package engine

import (
	"log"
	"os/exec"

	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
)

// Generate will run aks-engine generate on a given cluster definition
func (e *Engine) Generate() error {
	cmd := exec.Command("./bin/aks-engine", "generate", e.Config.ClusterDefinitionTemplate, "--output-directory", e.Config.GeneratedDefinitionPath)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to generate aks-engine template with cluster definition - %s: %s\n", e.Config.ClusterDefinitionTemplate, err)
		log.Printf("Command:./bin/aks-engine generate %s --output-directory %s\n", e.Config.ClusterDefinitionTemplate, e.Config.GeneratedDefinitionPath)
		log.Printf("Output:%s\n", out)
		return err
	}
	return nil
}

// Deploy will run aks-engine deploy on a given cluster definition
func (e *Engine) Deploy(location string) error {
	cmd := exec.Command("./bin/aks-engine", "deploy",
		"--location", location,
		"--api-model", e.Config.ClusterDefinitionPath,
		"--dns-prefix", e.Config.DefinitionName,
		"--output-directory", e.Config.GeneratedDefinitionPath,
		"--resource-group", e.Config.DefinitionName,
	)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to deploy aks-engine template with cluster definition - %s: %s\n", e.Config.ClusterDefinitionTemplate, err)
		log.Printf("Output:%s\n", out)
		return err
	}
	return nil
}
