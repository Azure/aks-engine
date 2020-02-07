//+build test
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package azure

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/Azure/aks-engine/test/e2e/engine"
	"github.com/Azure/aks-engine/test/e2e/kubernetes/util"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-08-01/network"

	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

// Storage provides access to StorageAccount objects
type Storage interface {
	CreateStorageAccount() error
	SetConnectionString() error
	CreateFileShare(name string) error
	UploadFiles(source, destination string) error
	DownloadFiles(source, destination string) error
	DeleteFiles(source string) error
}

// Account represents an Azure account
type Account struct {
	User            *User  `json:"user"`
	TenantID        string `json:"tenantId" envconfig:"TENANT_ID" required:"true"`
	SubscriptionID  string `json:"id" envconfig:"SUBSCRIPTION_ID" required:"true"`
	ResourceGroup   ResourceGroup
	Deployment      Deployment
	StorageAccount  *StorageAccount
	TimeoutCommands bool
}

// ResourceGroup represents a collection of azure resources
type ResourceGroup struct {
	Name     string            `json:"name"`
	Location string            `json:"location"`
	Tags     map[string]string `json:"tags"`
}

// VM represents an azure vm
type VM struct {
	Name string `json:"name"`
}

// Deployment represents a deployment of an acs cluster
type Deployment struct {
	Name              string // Name of the deployment
	TemplateDirectory string // engine.GeneratedDefinitionPath
}

// StorageAccount represents an azure storage account
type StorageAccount struct {
	Name             string
	ConnectionString string `json:"connectionString"`
	ResourceGroup    ResourceGroup
	TimeoutCommands  bool
}

// User represents the user currently logged into an Account
type User struct {
	ID     string `json:"name" envconfig:"CLIENT_ID" required:"true"`
	Secret string `envconfig:"CLIENT_SECRET" required:"true"`
	Type   string `json:"type"`
}

// NewAccount will parse env vars and return a new account struct
func NewAccount() (*Account, error) {
	a := new(Account)
	if err := envconfig.Process("account", a); err != nil {
		return nil, err
	}
	u := new(User)
	if err := envconfig.Process("user", u); err != nil {
		return nil, err
	}
	a.User = u
	a.StorageAccount = new(StorageAccount)

	cmd := exec.Command("which", "timeout")
	util.PrintCommand(cmd)
	_, err := cmd.CombinedOutput()
	if err == nil {
		a.TimeoutCommands = true
		a.StorageAccount.TimeoutCommands = true
	}
	return a, nil
}

// Login will login to a given subscription
func (a *Account) Login() error {
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "login",
			"--service-principal",
			"--username", a.User.ID,
			"--password", a.User.Secret,
			"--tenant", a.TenantID)
	} else {
		cmd = exec.Command("az", "login",
			"--service-principal",
			"--username", a.User.ID,
			"--password", a.User.Secret,
			"--tenant", a.TenantID)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("output:%s\n", out)
		return err
	}
	return nil
}

// SetSubscription will call az account set --subscription for the given Account
func (a *Account) SetSubscription() error {
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "account", "set", "--subscription", a.SubscriptionID)
	} else {
		cmd = exec.Command("az", "account", "set", "--subscription", a.SubscriptionID)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to set subscription (%s):%s", a.SubscriptionID, err)
		log.Printf("Output:%s\n", out)
		return err
	}
	return nil
}

// CreateGroupWithRetry invokes CreateGroup, retrying up to a timeout
func (a *Account) CreateGroupWithRetry(name, location string, sleep, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan error)
	var mostRecentCreateGroupWithRetryError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- a.CreateGroup(name, location):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case mostRecentCreateGroupWithRetryError = <-ch:
			if mostRecentCreateGroupWithRetryError == nil {
				return nil
			}
		case <-ctx.Done():
			return errors.Errorf("CreateGroupWithRetry timed out: %s\n", mostRecentCreateGroupWithRetryError)
		}
	}
}

// CreateGroup will create a resource group in a given location
//--tags "type=${RESOURCE_GROUP_TAG_TYPE:-}" "now=$(date +%s)" "job=${JOB_BASE_NAME:-}" "buildno=${BUILD_NUM:-}"
func (a *Account) CreateGroup(name, location string) error {
	now := fmt.Sprintf("now=%v", time.Now().Unix())
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "group", "create", "--name", name, "--location", location, "--tags", now)
	} else {
		cmd = exec.Command("az", "group", "create", "--name", name, "--location", location, "--tags", now)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to create resource group (%s) in %s:%s", name, location, err)
		log.Printf("Output:%s\n", out)
		return err
	}
	r := ResourceGroup{
		Name:     name,
		Location: location,
		Tags: map[string]string{
			"now": now,
		},
	}
	a.ResourceGroup = r
	return nil
}

// ShowGroupWithRetry invokes ShowGroup, retrying up to a timeout
func (a *Account) ShowGroupWithRetry(name string, sleep, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan error)
	var mostRecentShowGroupWithRetryError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- a.ShowGroup(name):
				time.Sleep(sleep)
			}
		}
	}()
	for {
		select {
		case mostRecentShowGroupWithRetryError = <-ch:
			if mostRecentShowGroupWithRetryError == nil {
				return nil
			}
		case <-ctx.Done():
			return errors.Errorf("ShowGroupWithRetry timed out: %s\n", mostRecentShowGroupWithRetryError)
		}
	}
}

// ShowGroup will validate that a resource group exists
func (a *Account) ShowGroup(name string) error {
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "group", "show", "-n", name)
	} else {
		cmd = exec.Command("az", "group", "show", "-n", name)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to show resource group (%s): %s", name, err)
		log.Printf("Output:%s\n", out)
		return err
	}
	return nil
}

// DeleteGroup deletes a given resource group by name
func (a *Account) DeleteGroup(name string, wait bool) error {
	var cmd *exec.Cmd
	if !wait {
		if a.TimeoutCommands {
			cmd = exec.Command("timeout", "60", "az", "group", "delete", "--name", name, "--no-wait", "--yes")
		} else {
			cmd = exec.Command("az", "group", "delete", "--name", name, "--no-wait", "--yes")
		}
	} else {
		cmd = exec.Command("az", "group", "delete", "--name", name, "--yes")
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to delete resource group (%s):%s", name, out)
		return err
	}
	return nil
}

// CreateDeployment will deploy a cluster to a given resource group using the template and parameters on disk
func (a *Account) CreateDeployment(name string, e *engine.Engine) error {
	d := Deployment{
		Name:              name,
		TemplateDirectory: e.Config.GeneratedDefinitionPath,
	}

	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Print(".")
			case <-quit:
				fmt.Print("\n")
				ticker.Stop()
				return
			}
		}
	}()

	cmd := exec.Command("az", "group", "deployment", "create",
		"--name", d.Name,
		"--resource-group", a.ResourceGroup.Name,
		"--template-file", e.Config.GeneratedTemplatePath,
		"--parameters", e.Config.GeneratedParametersPath)
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("\nError from deployment for %s in resource group %s:%s\n", d.Name, a.ResourceGroup.Name, err)
		log.Printf("Command Output: %s\n", out)
		return err
	}
	quit <- true
	a.Deployment = d
	return nil
}

// CreateVnet will create a vnet in a resource group
func (a *Account) CreateVnet(vnet, addressPrefixes string) error {
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "network", "vnet", "create", "-g",
			a.ResourceGroup.Name, "-n", vnet, "--address-prefixes", addressPrefixes)
	} else {
		cmd = exec.Command("az", "network", "vnet", "create", "-g",
			a.ResourceGroup.Name, "-n", vnet, "--address-prefixes", addressPrefixes)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to create vnet with the following command:\n az network vnet create -g %s -n %s --address-prefixes %s\n Output:%s\n", a.ResourceGroup.Name, vnet, addressPrefixes, out)
		return err
	}
	return nil
}

// ListRGRouteTableResult defines a struct for making a multi-value channel result type
type ListRGRouteTableResult struct {
	routeTables []network.RouteTable
	err         error
}

// ListRGRouteTable will list the route table for a resource group
func (a *Account) ListRGRouteTable() ListRGRouteTableResult {
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "network", "route-table", "list", "-g",
			a.ResourceGroup.Name, "-o", "json")
	} else {
		cmd = exec.Command("az", "network", "route-table", "list", "-g",
			a.ResourceGroup.Name, "-o", "json")
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to list route table with the following command:\n az network route-table list -g %s -o json | jq -r '.[].id'\n Output:%s\n", a.ResourceGroup.Name, out)
		return ListRGRouteTableResult{
			routeTables: nil,
			err:         err,
		}
	}
	routeTables := []network.RouteTable{}
	err = json.Unmarshal(out, &routeTables)
	if err != nil {
		log.Printf("Error unmarshalling route tables json:%s\n", err)
		return ListRGRouteTableResult{
			routeTables: nil,
			err:         err,
		}
	}
	return ListRGRouteTableResult{
		routeTables: routeTables,
		err:         err,
	}
}

// AddSubnetsToRouteTable will add the subnets in this cluster config to the VNET route table
func (a *Account) AddSubnetsToRouteTable(routeTableID, vnetName string, subnets []string) error {
	for _, subnet := range subnets {
		var cmd *exec.Cmd
		if a.TimeoutCommands {
			cmd = exec.Command("timeout", "60", "az", "network", "vnet", "subnet", "update", "-n", subnet, "-g",
				a.ResourceGroup.Name, "--vnet-name", vnetName, "--route-table", routeTableID)
		} else {
			cmd = exec.Command("az", "network", "vnet", "subnet", "update", "-n", subnet, "-g",
				a.ResourceGroup.Name, "--vnet-name", vnetName, "--route-table", routeTableID)
		}
		util.PrintCommand(cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Printf("Error while trying to add route table to subnet: %s\n", out)
			return err
		}
	}
	return nil
}

// GetRGRouteTable continuously calls the Azure API until we get a route table result
func (a *Account) GetRGRouteTable(timeout time.Duration) (network.RouteTable, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ch := make(chan ListRGRouteTableResult)
	var mostRecentListRGRouteTableError error
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- a.ListRGRouteTable():
				time.Sleep(5 * time.Second)
			}
		}
	}()
	for {
		select {
		case result := <-ch:
			mostRecentListRGRouteTableError = result.err
			if result.err == nil && result.routeTables != nil && len(result.routeTables) == 1 {
				return result.routeTables[0], nil
			}
		case <-ctx.Done():
			return network.RouteTable{}, mostRecentListRGRouteTableError
		}
	}
}

// CreateSubnet will create a subnet in a vnet in a resource group
func (a *Account) CreateSubnet(vnet, subnetName, subnetPrefix string) error {
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "network", "vnet", "subnet", "create", "-g",
			a.ResourceGroup.Name, "--vnet-name", vnet, "--name", subnetName, "--address-prefix", subnetPrefix)
	} else {
		cmd = exec.Command("az", "network", "vnet", "subnet", "create", "-g",
			a.ResourceGroup.Name, "--vnet-name", vnet, "--name", subnetName, "--address-prefix", subnetPrefix)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to create subnet in vnet with the following command:\n az network vnet subnet create -g %s --vnet-name %s --name %s --address-prefix %s \n Output:%s\n", a.ResourceGroup.Name, vnet, subnetName, subnetPrefix, out)
		return err
	}
	return nil
}

// RouteTable holds information from running az network route-table list
type RouteTable struct {
	ID                string `json:"id"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ProvisioningState string `json:"provisioningState"`
	ResourceGroup     string `json:"resourceGroup"`
}

// UpdateRouteTables is used to updated a vnet with the appropriate route tables
func (a *Account) UpdateRouteTables(subnet, vnet string) error {
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "network", "route-table", "list", "-g", a.ResourceGroup.Name)
	} else {
		cmd = exec.Command("az", "network", "route-table", "list", "-g", a.ResourceGroup.Name)
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to get route table list!\n Output:%s\n", out)
		return err
	}
	rts := []RouteTable{}
	json.Unmarshal(out, &rts)

	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "network", "vnet", "subnet", "update",
			"-n", subnet, "-g", a.ResourceGroup.Name, "--vnet-name", vnet, "--route-table", rts[0].Name)
	} else {
		cmd = exec.Command("az", "network", "vnet", "subnet", "update",
			"-n", subnet, "-g", a.ResourceGroup.Name, "--vnet-name", vnet, "--route-table", rts[0].Name)
	}
	util.PrintCommand(cmd)
	out, err = cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to update vnet route tables:%s\n", out)
		return err
	}
	return nil
}

// GetHosts will get a list of vms in the resource group
func (a *Account) GetHosts(name string) ([]VM, error) {
	v := []VM{{}}
	var resourceGroup string
	if name != "" {
		resourceGroup = name
	} else {
		resourceGroup = a.ResourceGroup.Name
	}
	err := a.ShowGroupWithRetry(resourceGroup, 3*time.Second, 10*time.Second)
	if err != nil {
		log.Printf("Unabled to validate that resource group %s already exists\n", resourceGroup)
		return v, nil
	}
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "vm", "list", "-g", resourceGroup)
	} else {
		cmd = exec.Command("az", "vm", "list", "-g", resourceGroup)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to get vm list:%s\n", out)
		return nil, err
	}
	err = json.Unmarshal(out, &v)
	if err != nil {
		log.Printf("Error unmarshalling VM json:%s\n", err)
		log.Printf("JSON:%s\n", out)
		return nil, err
	}
	return v, nil
}

// SetResourceGroup will set the account resource group
func (a *Account) SetResourceGroup(name string) error {
	if a.ResourceGroup.Name != "" {
		return nil
	}
	var cmd *exec.Cmd
	if a.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "group", "show", "-g", name)
	} else {
		cmd = exec.Command("az", "group", "show", "-g", name)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to show resource group:%s\n", out)
		return err
	}
	if len(out) == 0 {
		log.Printf("Resource group %s does not exist\n", name)
		return errors.New("Resource group not found")
	}
	a.ResourceGroup = ResourceGroup{}
	err = json.Unmarshal(out, &a.ResourceGroup)
	if err != nil {
		log.Printf("Error unmarshalling resource group json:%s\n", err)
		log.Printf("JSON:%s\n", out)
		return err
	}
	return nil
}

// IsClusterExpired will return true if a deployment was created more than t nanoseconds ago, or if timestamp is not found
func (a *Account) IsClusterExpired(d time.Duration) bool {
	tag, err := strconv.ParseInt(a.ResourceGroup.Tags["now"], 10, 64)
	if err != nil {
		log.Printf("Error parsing RG now tag:%s\n", err)
		return true
	}
	t := time.Unix(tag, 0)
	age := time.Since(t)
	log.Printf("Cluster is %v hours old\n", int(age.Hours()))
	return age > d
}

// FetchActivityLog gets all the failures from the activity log for the provided resource group.
func (a *Account) FetchActivityLog(rg string) (string, error) {
	var cmd = exec.Command("az", "monitor", "activity-log", "list", "--resource-group", rg, "--status", "Failed")
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// CreateStorageAccount will create a new Azure Storage Account
func (sa *StorageAccount) CreateStorageAccount() error {
	var cmd *exec.Cmd
	if sa.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "storage", "account", "create", "--name", sa.Name, "--resource-group", sa.ResourceGroup.Name)
	} else {
		cmd = exec.Command("az", "storage", "account", "create", "--name", sa.Name, "--resource-group", sa.ResourceGroup.Name)
	}
	util.PrintCommand(cmd)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to create storage account: %s", out)
		return err
	}
	return nil
}

// SetConnectionString will set the storage account connection string
func (sa *StorageAccount) SetConnectionString() error {
	var cmd *exec.Cmd
	if sa.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "storage", "account", "show-connection-string", "-g", sa.ResourceGroup.Name, "-n", sa.Name)
	} else {
		cmd = exec.Command("az", "storage", "account", "show-connection-string", "-g", sa.ResourceGroup.Name, "-n", sa.Name)
	}
	log.Printf("Getting connection string for storage account %s...\n", sa.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to get connection-string:%s\n", out)
		return err
	}
	err = json.Unmarshal(out, &sa)
	if err != nil {
		log.Printf("Error unmarshalling account json:%s\n", err)
		log.Printf("JSON:%s\n", out)
		return err
	}
	return nil
}

// CreateFileShare will create a file share in a storage account if it doesn't already exist
func (sa *StorageAccount) CreateFileShare(name string) error {
	var cmd *exec.Cmd
	if sa.TimeoutCommands {
		cmd = exec.Command("timeout", "60", "az", "storage", "share", "create", "--name", name, "--account-name", sa.Name, "--connection-string", sa.ConnectionString)
	} else {
		cmd = exec.Command("az", "storage", "share", "create", "--name", name, "--account-name", sa.Name, "--connection-string", sa.ConnectionString)
	}
	log.Printf("Creating file share %s in storage account %s...\n", name, sa.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to create file share: %s", out)
		return err
	}
	return nil
}

// UploadFiles will upload the output directory to storage
func (sa *StorageAccount) UploadFiles(source, destination string) error {
	cmd := exec.Command("az", "storage", "file", "upload-batch", "--destination", destination, "--source", source, "--account-name", sa.Name, "--connection-string", sa.ConnectionString)
	log.Printf("Uploading files to %s in storage account %s...\n", destination, sa.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying upload files to file share:%s\n", out)
		return err
	}
	return nil
}

// DownloadFiles will download the output directory from storage
func (sa *StorageAccount) DownloadFiles(source, destination string) error {
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		os.Mkdir(destination, os.ModeDir)
	}
	cmd := exec.Command("az", "storage", "file", "download-batch", "--destination", destination, "--source", source, "--account-name", sa.Name, "--connection-string", sa.ConnectionString)
	log.Printf("Downloading files from %s in storage account %s...\n", source, sa.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying download files from %s in storage account %s: %s\n", source, sa.Name, out)
		return err
	}
	return nil
}

// DeleteFiles deletes files from an Azure storage file share
func (sa *StorageAccount) DeleteFiles(source string) error {
	cmd := exec.Command("az", "storage", "file", "delete-batch", "--source", source, "--account-name", sa.Name, "--connection-string", sa.ConnectionString)
	log.Printf("Deleting files in %s in storage account %s...\n", source, sa.Name)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error while trying to delete files from %s: %s", source, out)
		return err
	}
	return nil
}
