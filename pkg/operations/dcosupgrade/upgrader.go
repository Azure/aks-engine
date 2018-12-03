// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package dcosupgrade

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Azure/aks-engine/pkg/operations"
	"github.com/pkg/errors"
)

type nodeHealthReport struct {
	IP     string `json:"host_ip"`
	Role   string `json:"role"`
	Health int    `json:"health"`
}

type healthReport struct {
	Nodes []nodeHealthReport `json:"nodes"`
}

var bootstrapUpgradeScript = `#!/bin/bash

echo "Starting upgrade configuration"
if [ ! -e /opt/azure/dcos/upgrade/NEW_VERSION/upgrade_url ]; then
  echo "Setting up bootstrap node"
  rm -rf /opt/azure/dcos/upgrade/NEW_VERSION
  mkdir -p /opt/azure/dcos/upgrade/NEW_VERSION/genconf
  cp /opt/azure/dcos/genconf/config.yaml /opt/azure/dcos/genconf/ip-detect /opt/azure/dcos/upgrade/NEW_VERSION/genconf/
  cd /opt/azure/dcos/upgrade/NEW_VERSION/
  curl -s -O https://dcos-mirror.azureedge.net/dcos-NEW_DASHED_VERSION/dcos_generate_config.sh
  bash dcos_generate_config.sh --generate-node-upgrade-script CURR_VERSION | tee /opt/azure/dcos/upgrade/NEW_VERSION/log
  process=\$(docker ps -f ancestor=nginx -q)
  if [ ! -z "\$process" ]; then
    echo "Stopping nginx service \$process"
    docker kill \$process
  fi
  echo "Starting nginx service \$process"
  docker run -d -p 8086:80 -v \$PWD/genconf/serve:/usr/share/nginx/html:ro nginx
  docker ps
  grep 'Node upgrade script URL' /opt/azure/dcos/upgrade/NEW_VERSION/log | awk -F ': ' '{print \$2}' | cat > /opt/azure/dcos/upgrade/NEW_VERSION/upgrade_url
fi
upgrade_url=\$(cat /opt/azure/dcos/upgrade/NEW_VERSION/upgrade_url)
if [ -z \${upgrade_url} ]; then
  rm -f /opt/azure/dcos/upgrade/NEW_VERSION/upgrade_url
  echo "Failed to set up bootstrap node. Please try again"
  exit 1
else
  echo "Setting up bootstrap node completed. Node upgrade script URL \${upgrade_url}"
fi
`

var nodeUpgradeScript = `#!/bin/bash

echo "Starting node upgrade"
mkdir -p /opt/azure/dcos/upgrade/NEW_VERSION
cd /opt/azure/dcos/upgrade/NEW_VERSION
curl -O UPGRADE_SCRIPT_URL
bash ./dcos_node_upgrade.sh

`

func (uc *UpgradeCluster) runUpgrade() error {
	if uc.ClusterTopology.DataModel.Properties.OrchestratorProfile.DcosConfig == nil ||
		uc.ClusterTopology.DataModel.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile == nil {
		return errors.New("BootstrapProfile is not set")
	}

	cs := uc.ClusterTopology.DataModel
	newVersion := cs.Properties.OrchestratorProfile.OrchestratorVersion
	dashedVersion := strings.Replace(newVersion, ".", "-", -1)

	masterDNS := cs.GetAzureProdFQDN()
	masterCount := cs.Properties.MasterProfile.Count
	bootstrapIP := cs.Properties.OrchestratorProfile.DcosConfig.BootstrapProfile.StaticIP
	uc.Logger.Infof("masterDNS:%s masterCount:%d bootstrapIP:%s", masterDNS, masterCount, bootstrapIP)

	// copy SSH key to master
	uc.Logger.Infof("Copy SSH key to master")
	out, err := operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, fmt.Sprintf("cat << END > .ssh/id_rsa_cluster\n%s\nEND\n", string(uc.SSHKey)))
	if err != nil {
		uc.Logger.Errorf(out)
		return err
	}
	// set SSH key permissions
	out, err = operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, "chmod 600 .ssh/id_rsa_cluster")
	if err != nil {
		uc.Logger.Errorf(out)
		return err
	}
	// upgrade bootstrap node
	bootstrapScript := strings.Replace(bootstrapUpgradeScript, "CURR_VERSION", uc.CurrentDcosVersion, -1)
	bootstrapScript = strings.Replace(bootstrapScript, "NEW_VERSION", newVersion, -1)
	bootstrapScript = strings.Replace(bootstrapScript, "NEW_DASHED_VERSION", dashedVersion, -1)

	upgradeScriptURL, err := uc.upgradeBootstrapNode(masterDNS, bootstrapIP, bootstrapScript)
	if err != nil {
		return err
	}

	nodeScript := strings.Replace(nodeUpgradeScript, "NEW_VERSION", newVersion, -1)
	nodeScript = strings.Replace(nodeScript, "UPGRADE_SCRIPT_URL", upgradeScriptURL, -1)

	// upgrade master nodes
	if err = uc.upgradeMasterNodes(masterDNS, masterCount, nodeScript); err != nil {
		return err
	}

	// get the node list
	out, err = operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, "curl -s http://localhost:1050/system/health/v1/nodes")
	if err != nil {
		uc.Logger.Errorf(out)
		return err
	}
	uc.Logger.Info(out)
	nodes := &healthReport{}
	if err = json.Unmarshal([]byte(out), nodes); err != nil {
		return err
	}
	// upgrade agent nodes
	return uc.upgradeAgentNodes(masterDNS, nodes)
}

func (uc *UpgradeCluster) upgradeBootstrapNode(masterDNS, bootstrapIP, bootstrapScript string) (string, error) {
	// copy bootstrap script to master
	uc.Logger.Infof("Copy bootstrap script to master")
	out, err := operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, fmt.Sprintf("cat << END > bootstrap_upgrade.sh\n%s\nEND\n", bootstrapScript))
	if err != nil {
		uc.Logger.Errorf(out)
		return "", err
	}
	// set script permissions
	out, err = operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, "chmod 755 ./bootstrap_upgrade.sh")
	if err != nil {
		uc.Logger.Errorf(out)
		return "", err
	}
	// copy bootstrap script to the bootstrap node
	uc.Logger.Infof("Copy bootstrap script to the bootstrap node")
	cmd := fmt.Sprintf("scp -i .ssh/id_rsa_cluster -o ConnectTimeout=30 -o StrictHostKeyChecking=no bootstrap_upgrade.sh %s:", bootstrapIP)
	out, err = operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, cmd)
	if err != nil {
		uc.Logger.Errorf(out)
		return "", err
	}
	// run bootstrap script
	uc.Logger.Infof("Run bootstrap upgrade script")
	cmd = fmt.Sprintf("ssh -i .ssh/id_rsa_cluster -o ConnectTimeout=30 -o StrictHostKeyChecking=no %s sudo ./bootstrap_upgrade.sh", bootstrapIP)
	out, err = operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, cmd)
	if err != nil {
		uc.Logger.Errorf(out)
		return "", err
	}
	uc.Logger.Info(out)
	// retrieve upgrade script URL
	var url string
	arr := strings.Split(out, "\n")
	prefix := "Setting up bootstrap node completed. Node upgrade script URL"
	for _, str := range arr {
		if strings.HasPrefix(str, prefix) {
			url = strings.TrimSpace(str[len(prefix):])
			break
		}
	}
	if len(url) == 0 {
		return "", errors.New("Undefined upgrade script URL")
	}
	return url, nil
}

func (uc *UpgradeCluster) upgradeMasterNodes(masterDNS string, masterCount int, nodeScript string) error {
	// run master upgrade script
	catCmd := fmt.Sprintf("cat << END > node_upgrade.sh\n%s\nEND\n", nodeScript)
	for i := 0; i < masterCount; i++ {
		uc.Logger.Infof("Upgrading master node #%d", i+1)
		// check current version
		out, err := operations.RemoteRun("azureuser", masterDNS, 2200+i, uc.SSHKey, "grep version /opt/mesosphere/etc/dcos-version.json | cut -d '\"' -f 4")
		if err != nil {
			uc.Logger.Errorf(out)
			return err
		}
		if strings.TrimSpace(out) == uc.ClusterTopology.DataModel.Properties.OrchestratorProfile.OrchestratorVersion {
			uc.Logger.Infof("Master node is up-to-date. Skipping upgrade")
			continue
		}
		// copy script to the node
		out, err = operations.RemoteRun("azureuser", masterDNS, 2200+i, uc.SSHKey, catCmd)
		if err != nil {
			uc.Logger.Errorf(out)
			return err
		}
		// set script permissions
		out, err = operations.RemoteRun("azureuser", masterDNS, 2200+i, uc.SSHKey, "chmod 755 ./node_upgrade.sh")
		if err != nil {
			uc.Logger.Errorf(out)
			return err
		}
		// run the script
		out, err = operations.RemoteRun("azureuser", masterDNS, 2200+i, uc.SSHKey, "sudo ./node_upgrade.sh")
		if err != nil {
			uc.Logger.Errorf(out)
			return err
		}
		uc.Logger.Info(out)
	}
	return nil
}

func (uc *UpgradeCluster) upgradeAgentNodes(masterDNS string, nodes *healthReport) error {
	for _, node := range nodes.Nodes {
		if node.Role == "master" {
			continue
		}
		uc.Logger.Infof("Upgrading %s %s", node.Role, node.IP)

		// check current version
		cmd := fmt.Sprintf("ssh -i .ssh/id_rsa_cluster -o ConnectTimeout=30 -o StrictHostKeyChecking=no %s grep version /opt/mesosphere/etc/dcos-version.json | cut -d '\"' -f 4", node.IP)
		out, err := operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, cmd)
		if err != nil {
			uc.Logger.Errorf(out)
			return err
		}
		if strings.TrimSpace(out) == uc.ClusterTopology.DataModel.Properties.OrchestratorProfile.OrchestratorVersion {
			uc.Logger.Infof("Agent node is up-to-date. Skipping upgrade")
			continue
		}
		// copy script to the node
		cmd = fmt.Sprintf("scp -i .ssh/id_rsa_cluster -o ConnectTimeout=30 -o StrictHostKeyChecking=no node_upgrade.sh %s:", node.IP)
		out, err = operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, cmd)
		if err != nil {
			uc.Logger.Errorf(out)
			return err
		}
		// run the script
		cmd = fmt.Sprintf("ssh -i .ssh/id_rsa_cluster -o ConnectTimeout=30 -o StrictHostKeyChecking=no %s sudo ./node_upgrade.sh", node.IP)
		out, err = operations.RemoteRun("azureuser", masterDNS, 2200, uc.SSHKey, cmd)
		if err != nil {
			uc.Logger.Errorf(out)
			return err
		}
		uc.Logger.Info(out)
	}
	return nil
}
