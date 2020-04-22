// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package transform

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-12-01/compute"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	//Field names
	customDataFieldName               = "customData"
	dependsOnFieldName                = "dependsOn"
	hardwareProfileFieldName          = "hardwareProfile"
	imageReferenceFieldName           = "imageReference"
	nameFieldName                     = "name"
	osProfileFieldName                = "osProfile"
	propertiesFieldName               = "properties"
	resourcesFieldName                = "resources"
	storageProfileFieldName           = "storageProfile"
	typeFieldName                     = "type"
	vmSizeFieldName                   = "vmSize"
	dataDisksFieldName                = "dataDisks"
	createOptionFieldName             = "createOption"
	tagsFieldName                     = "tags"
	managedDiskFieldName              = "managedDisk"
	windowsConfigurationFieldName     = "windowsConfiguration"
	platformFaultDomainCountFieldName = "platformFaultDomainCount"
	singlePlacementGroupFieldName     = "singlePlacementGroup"
	proximityPlacementGroupFieldName  = "proximityPlacementGroup"

	// ARM resource Types
	nsgResourceType  = "Microsoft.Network/networkSecurityGroups"
	rtResourceType   = "Microsoft.Network/routeTables"
	vmResourceType   = "Microsoft.Compute/virtualMachines"
	vmExtensionType  = "Microsoft.Compute/virtualMachines/extensions"
	nicResourceType  = "Microsoft.Network/networkInterfaces"
	vnetResourceType = "Microsoft.Network/virtualNetworks"
	vmasResourceType = "Microsoft.Compute/availabilitySets"
	vmssResourceType = "Microsoft.Compute/virtualMachineScaleSets"
	lbResourceType   = "Microsoft.Network/loadBalancers"

	// resource ids
	nsgID     = "nsgID"
	rtID      = "routeTableID"
	vnetID    = "vnetID"
	agentLbID = "agentLbID"
)

// Translator defines all required interfaces for i18n.Translator.
type Translator interface {
	// T translates a text string, based on GNU's gettext library.
	T(msgid string, vars ...interface{}) string
	// NT translates a text string into the appropriate plural form, based on GNU's gettext library.
	NT(msgid, msgidPlural string, n int, vars ...interface{}) string
	// Errorf produces an error with a translated error string.
	Errorf(msgid string, vars ...interface{}) error
	// NErrorf produces an error with a translated error string in the appropriate plural form.
	NErrorf(msgid, msgidPlural string, n int, vars ...interface{}) error
}

// Transformer represents the object that transforms template
type Transformer struct {
	Translator Translator
}

type tMap map[string]interface{}
type resource map[string]interface{}

func (t tMap) Resources(logger *logrus.Entry) []resource {
	resourcesInterfaces := t[resourcesFieldName].([]interface{})
	resources := make([]resource, 0)
	for index, ri := range resourcesInterfaces {
		if r, ok := ri.(map[string]interface{}); ok {
			resources = append(resources, r)
		} else {
			logger.Warnf("Template improperly formatted for resource at index %d", index)
		}
	}
	return resources
}

func (r resource) Type() string {
	return r[typeFieldName].(string)
}

func (r resource) Name() string {
	return r[nameFieldName].(string)
}

func (r resource) Properties() map[string]interface{} {
	prop, ok := r[propertiesFieldName].(map[string]interface{})
	if !ok {
		return map[string]interface{}{}
	}
	return prop
}

func (r resource) RemoveProperty(logger *logrus.Entry, key string) {
	properties := r.Properties()
	_, ok := properties[key]
	if ok {
		logger.Infof("Removing %s property from %s", key, r.Name())
		delete(properties, key)
	}
}

func (r resource) removeDependencyType(logger *logrus.Entry, depType string) []string {
	deps := r.DependsOn()
	var newDependsOn []string
	for _, dep := range deps {
		depVal := dep.(string)
		if !strings.Contains(depVal, depType) {
			logger.Infof("Removing %s dependency from %s", depType, r.Name())
			newDependsOn = append(newDependsOn, depVal)
		}
	}
	return newDependsOn
}

func (r resource) DependsOn() []interface{} {
	deps, ok := r[dependsOnFieldName].([]interface{})
	if !ok {
		return []interface{}{}
	}
	return deps
}

func (t *Transformer) RemoveImmutableResourceProperties(logger *logrus.Entry, templateMap map[string]interface{}) {
	tm := tMap(templateMap)
	for _, resource := range tm.Resources(logger) {
		if resource.Type() == vmssResourceType {
			resource.RemoveProperty(logger, platformFaultDomainCountFieldName)
			resource.RemoveProperty(logger, singlePlacementGroupFieldName)
			resource.RemoveProperty(logger, proximityPlacementGroupFieldName)
		}
	}
}

func (t *Transformer) RemoveJumpboxResourcesFromTemplate(logger *logrus.Entry, templateMap map[string]interface{}) error {
	logger.Debugf("Running RemoveJumpboxResourcesFromTemplate...")
	resources := templateMap[resourcesFieldName].([]interface{})
	indexesToRemove := []int{}
	for index, resource := range resources {
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			return errors.Errorf("Template improperly formatted for resource")
		}

		resourceName, ok := resourceMap[nameFieldName].(string)
		if !ok {
			logger.Warnf("Resource does not have a name property")
			continue
		} else if strings.Contains(resourceName, "variables('jumpbox") || strings.Contains(resourceName, "parameters('jumpbox") {
			indexesToRemove = append(indexesToRemove, index)
		}
	}
	templateMap[resourcesFieldName] = removeIndexesFromArray(resources, indexesToRemove)
	return nil
}

// NormalizeForK8sSLBScalingOrUpgrade takes a template and removes elements that are unwanted in a K8s Standard LB cluster scale up/down case
func (t *Transformer) NormalizeForK8sSLBScalingOrUpgrade(logger *logrus.Entry, templateMap map[string]interface{}) error {
	logger.Debugf("Running NormalizeForK8sSLBScalingOrUpgrade...")
	lbIndex := -1
	resources := templateMap[resourcesFieldName].([]interface{})

	for index, resource := range resources {
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			logger.Warnf("Template improperly formatted for resource")
			continue
		}

		resourceType, ok := resourceMap[typeFieldName].(string)
		resourceName := resourceMap[nameFieldName].(string)

		// remove agentLB if found
		if ok && resourceType == lbResourceType && strings.Contains(resourceName, "variables('agentLbName')") {
			lbIndex = index
		}
		// remove agentLB from dependsOn if found
		dependencies, ok := resourceMap[dependsOnFieldName].([]interface{})
		if !ok {
			continue
		}

		for dIndex := len(dependencies) - 1; dIndex >= 0; dIndex-- {
			dependency := dependencies[dIndex].(string)
			if strings.Contains(dependency, lbResourceType) || strings.Contains(dependency, agentLbID) {
				dependencies = append(dependencies[:dIndex], dependencies[dIndex+1:]...)
			}
		}

		if len(dependencies) > 0 {
			resourceMap[dependsOnFieldName] = dependencies
		} else {
			delete(resourceMap, dependsOnFieldName)
		}
	}
	indexesToRemove := []int{}
	if lbIndex != -1 {
		indexesToRemove = append(indexesToRemove, lbIndex)
	}
	templateMap[resourcesFieldName] = removeIndexesFromArray(resources, indexesToRemove)
	return nil
}

// NormalizeForK8sVMASScalingUp takes a template and removes elements that are unwanted in a K8s VMAS scale up/down case
func (t *Transformer) NormalizeForK8sVMASScalingUp(logger *logrus.Entry, templateMap map[string]interface{}) error {
	if err := t.NormalizeMasterResourcesForScaling(logger, templateMap); err != nil {
		return err
	}
	rtIndex := -1
	nsgIndex := -1
	vnetIndex := -1
	vmasIndexes := make([]int, 0)

	resources := templateMap[resourcesFieldName].([]interface{})
	for index, resource := range resources {
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			logger.Warnf("Template improperly formatted for resource")
			continue
		}

		resourceType, ok := resourceMap[typeFieldName].(string)
		resourceName := resourceMap[nameFieldName].(string)

		if ok && resourceType == nsgResourceType && !strings.Contains(resourceName, "variables('jumpboxNetworkSecurityGroupName')") {

			if nsgIndex != -1 {
				err := t.Translator.Errorf("Found 2 resources with type %s in the template. There should only be 1", nsgResourceType)
				logger.Errorf(err.Error())
				return err
			}
			nsgIndex = index
		}
		if ok && resourceType == rtResourceType {
			if rtIndex != -1 {
				err := t.Translator.Errorf("Found 2 resources with type %s in the template. There should only be 1", rtResourceType)
				logger.Warnf(err.Error())
				return err
			}
			rtIndex = index
		}
		if ok && resourceType == vnetResourceType {
			if vnetIndex != -1 {
				err := t.Translator.Errorf("Found 2 resources with type %s in the template. There should only be 1", vnetResourceType)
				logger.Warnf(err.Error())
				return err
			}
			vnetIndex = index
		}
		if ok && resourceType == vmasResourceType {
			// All availability sets can be removed
			vmasIndexes = append(vmasIndexes, index)
		}

		dependencies, ok := resourceMap[dependsOnFieldName].([]interface{})
		if !ok {
			continue
		}

		for dIndex := len(dependencies) - 1; dIndex >= 0; dIndex-- {
			dependency := dependencies[dIndex].(string)
			if strings.Contains(dependency, nsgResourceType) || strings.Contains(dependency, nsgID) ||
				strings.Contains(dependency, rtResourceType) || strings.Contains(dependency, rtID) ||
				strings.Contains(dependency, vnetResourceType) || strings.Contains(dependency, vnetID) ||
				strings.Contains(dependency, vmasResourceType) {
				dependencies = append(dependencies[:dIndex], dependencies[dIndex+1:]...)
			}
		}

		if len(dependencies) > 0 {
			resourceMap[dependsOnFieldName] = dependencies
		} else {
			delete(resourceMap, dependsOnFieldName)
		}
	}

	indexesToRemove := []int{}

	if rtIndex == -1 {
		logger.Debugf("Found no resources with type %s in the template.", rtResourceType)
	} else {
		indexesToRemove = append(indexesToRemove, rtIndex)
	}

	if vnetIndex != -1 {
		indexesToRemove = append(indexesToRemove, vnetIndex)
	}

	if len(vmasIndexes) != 0 {
		indexesToRemove = append(indexesToRemove, vmasIndexes...)
	}
	if nsgIndex > 0 {
		indexesToRemove = append(indexesToRemove, nsgIndex)
	}

	templateMap[resourcesFieldName] = removeIndexesFromArray(resources, indexesToRemove)

	return nil
}

func removeIndexesFromArray(array []interface{}, indexes []int) []interface{} {
	sort.Sort(sort.Reverse(sort.IntSlice(indexes)))
	for _, index := range indexes {
		array = append(array[:index], array[index+1:]...)
	}
	return array
}

// NormalizeMasterResourcesForScaling takes a template and removes elements that are unwanted in any scale up/down case
func (t *Transformer) NormalizeMasterResourcesForScaling(logger *logrus.Entry, templateMap map[string]interface{}) error {
	resources := templateMap[resourcesFieldName].([]interface{})
	indexesToRemove := []int{}
	//update master nodes resources
	for index, resource := range resources {
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			logger.Warnf("Template improperly formatted")
			continue
		}

		resourceType, ok := resourceMap[typeFieldName].(string)
		if !ok || resourceType != vmResourceType {
			var resourceName string
			resourceName, ok = resourceMap[nameFieldName].(string)
			if !ok {
				logger.Warnf("Template improperly formatted")
				continue
			}
			if strings.Contains(resourceName, "variables('masterVMNamePrefix')") && resourceType == vmExtensionType {
				indexesToRemove = append(indexesToRemove, index)
			}
			continue
		}

		resourceName, ok := resourceMap[nameFieldName].(string)
		if !ok {
			logger.Warnf("Template improperly formatted")
			continue
		}

		// make sure this is only modifying the master vms
		if !strings.Contains(resourceName, "variables('masterVMNamePrefix')") {
			continue
		}

		resourceProperties, ok := resourceMap[propertiesFieldName].(map[string]interface{})
		if !ok {
			logger.Warnf("Template improperly formatted")
			continue
		}

		hardwareProfile, ok := resourceProperties[hardwareProfileFieldName].(map[string]interface{})
		if !ok {
			logger.Warnf("Template improperly formatted")
			continue
		}

		if hardwareProfile[vmSizeFieldName] != nil {
			delete(hardwareProfile, vmSizeFieldName)
		}

		if !t.removeCustomData(logger, resourceProperties) || !t.removeDataDisks(logger, resourceProperties) || !t.removeImageReference(logger, resourceProperties) {
			continue
		}
	}
	templateMap[resourcesFieldName] = removeIndexesFromArray(resources, indexesToRemove)

	return nil
}

func (t *Transformer) removeCustomData(logger *logrus.Entry, resourceProperties map[string]interface{}) bool {
	osProfile, ok := resourceProperties[osProfileFieldName].(map[string]interface{})
	if !ok {
		logger.Warnf("Template improperly formatted")
		return ok
	}

	if osProfile[customDataFieldName] != nil && osProfile[windowsConfigurationFieldName] == nil {
		delete(osProfile, customDataFieldName)
	}
	return ok
}

func (t *Transformer) removeDataDisks(logger *logrus.Entry, resourceProperties map[string]interface{}) bool {
	storageProfile, ok := resourceProperties[storageProfileFieldName].(map[string]interface{})
	if !ok {
		logger.Warnf("Template improperly formatted. Could not find: %s", storageProfileFieldName)
		return ok
	}

	if storageProfile[dataDisksFieldName] != nil {
		delete(storageProfile, dataDisksFieldName)
	}
	return ok
}

func (t *Transformer) removeImageReference(logger *logrus.Entry, resourceProperties map[string]interface{}) bool {
	storageProfile, ok := resourceProperties[storageProfileFieldName].(map[string]interface{})
	if !ok {
		logger.Warnf("Template improperly formatted. Could not find: %s", storageProfileFieldName)
		return ok
	}

	if storageProfile[imageReferenceFieldName] != nil {
		delete(storageProfile, imageReferenceFieldName)
	}
	return ok
}

// NormalizeResourcesForK8sMasterUpgrade takes a template and removes elements that are unwanted in any scale up/down case
func (t *Transformer) NormalizeResourcesForK8sMasterUpgrade(logger *logrus.Entry, templateMap map[string]interface{}, isMasterManagedDisk bool, agentPoolsToPreserve map[string]bool) error {
	resources := templateMap[resourcesFieldName].([]interface{})
	resourceTypeToProcess := map[string]bool{vmResourceType: true, vmExtensionType: true,
		nicResourceType: true, vnetResourceType: true, nsgResourceType: true,
		lbResourceType: true, vmssResourceType: true, vmasResourceType: true}
	logger.Infoln(fmt.Sprintf("Resource count before running NormalizeResourcesForK8sMasterUpgrade: %d", len(resources)))

	filteredResources := resources[:0]

	// remove agent nodes resources if needed and set dataDisk createOption to attach
	for _, resource := range resources {
		filteredResources = append(filteredResources, resource)
		resourceMap, ok := resource.(map[string]interface{})
		if !ok {
			logger.Warnf("Template improperly formatted for field name: %s", resourcesFieldName)
			continue
		}

		resourceType, ok := resourceMap[typeFieldName].(string)
		if !ok {
			continue
		}

		_, process := resourceTypeToProcess[resourceType]
		if !process {
			continue
		}

		filteredResources = removeVMAS(logger, filteredResources, resourceMap)

		resourceName, ok := resourceMap[nameFieldName].(string)
		if !ok {
			logger.Warnf("Template improperly formatted for field name: %s", nameFieldName)
			continue
		}

		if resourceType == vmssResourceType || resourceType == vnetResourceType {
			RemoveNsgDependency(logger, resourceName, resourceMap)
			continue
		}

		if resourceType == lbResourceType {
			if strings.Contains(resourceName, "variables('masterInternalLbName')") {
				RemoveNsgDependency(logger, resourceName, resourceMap)
				continue
			}
		}

		if resourceType == nicResourceType {
			RemoveNsgDependency(logger, resourceName, resourceMap)
			if strings.Contains(resourceName, "variables('masterVMNamePrefix')") {
				continue
			} else {
				// Remove agent NICs if upgrade master nodes
				if agentPoolsToPreserve == nil {
					logger.Infoln(fmt.Sprintf("Removing nic: %s from template", resourceName))
					if len(filteredResources) > 0 {
						filteredResources = filteredResources[:len(filteredResources)-1]
					}
				}
				continue
			}
		}

		if strings.EqualFold(resourceType, vmResourceType) &&
			strings.Contains(resourceName, "variables('masterVMNamePrefix')") {
			resourceProperties, ok := resourceMap[propertiesFieldName].(map[string]interface{})
			if !ok {
				logger.Warnf("Template improperly formatted for field name: %s, resource name: %s", propertiesFieldName, resourceName)
				continue
			}

			storageProfile, ok := resourceProperties[storageProfileFieldName].(map[string]interface{})
			if !ok {
				logger.Warnf("Template improperly formatted: %s", storageProfileFieldName)
				continue
			}

			dataDisks, ok := storageProfile[dataDisksFieldName].([]interface{})
			if !ok {
				logger.Warnf("Template improperly formatted for field name: %s, property name: %s", storageProfileFieldName, dataDisksFieldName)
				continue
			}

			dataDisk, ok := dataDisks[0].(map[string]interface{})
			if !ok {
				logger.Warnf("Template improperly formatted for field name: %s, there is no data disks defined", dataDisksFieldName)
				continue
			}

			dataDisk[createOptionFieldName] = "attach"

			if isMasterManagedDisk {
				managedDisk := compute.ManagedDiskParameters{}
				id := "[concat('/subscriptions/', variables('subscriptionId'), '/resourceGroups/', variables('resourceGroup'),'/providers/Microsoft.Compute/disks/', variables('masterVMNamePrefix'), copyIndex(variables('masterOffset')),'-etcddisk')]"
				managedDisk.ID = &id
				diskInterface := &managedDisk
				dataDisk[managedDiskFieldName] = diskInterface
			}
		}

		tags, _ := resourceMap[tagsFieldName].(map[string]interface{})
		poolName := fmt.Sprint(tags["poolName"]) // poolName tag exists on agents only

		if resourceType == vmResourceType {
			logger.Infoln(fmt.Sprintf("Evaluating if agent pool: %s, resource: %s needs to be removed", poolName, resourceName))
			// Not an agent (could be a master VM)
			if tags["poolName"] == nil || strings.Contains(resourceName, "variables('masterVMNamePrefix')") {
				continue
			}

			logger.Infoln(fmt.Sprintf("agentPoolsToPreserve: %v...", agentPoolsToPreserve))

			if len(agentPoolsToPreserve) == 0 || !agentPoolsToPreserve[poolName] {
				logger.Infoln(fmt.Sprintf("Removing agent pool: %s, resource: %s from template", poolName, resourceName))
				if len(filteredResources) > 0 {
					filteredResources = filteredResources[:len(filteredResources)-1]
				}
			}
		} else if resourceType == vmExtensionType {
			logger.Infoln(fmt.Sprintf("Evaluating if extension: %s needs to be removed", resourceName))
			if strings.Contains(resourceName, "variables('masterVMNamePrefix')") {
				continue
			}

			logger.Infoln(fmt.Sprintf("agentPoolsToPreserve: %v...", agentPoolsToPreserve))

			removeExtension := true
			for poolName, preserve := range agentPoolsToPreserve {
				if strings.Contains(resourceName, "variables('"+poolName) && preserve {
					removeExtension = false
				}
			}

			if removeExtension {
				logger.Infoln(fmt.Sprintf("Removing extension: %s from template", resourceName))
				if len(filteredResources) > 0 {
					filteredResources = filteredResources[:len(filteredResources)-1]
				}
			}
		} else if resourceType == nsgResourceType {
			logger.Infoln(fmt.Sprintf("Removing nsg resource: %s from template", resourceName))
			{
				if len(filteredResources) > 0 {
					filteredResources = filteredResources[:len(filteredResources)-1]
				}
			}
		}
	}

	templateMap[resourcesFieldName] = filteredResources

	logger.Infoln(fmt.Sprintf("Resource count after running NormalizeResourcesForK8sMasterUpgrade: %d",
		len(templateMap[resourcesFieldName].([]interface{}))))
	return nil
}

func removeVMAS(logger *logrus.Entry, resources []interface{}, resource resource) []interface{} {
	// remove vmas
	if strings.EqualFold(resource.Type(), vmasResourceType) {
		return resources[:len(resources)-1]
	}
	// remove dependencies on vmas
	if strings.EqualFold(resource.Type(), vmResourceType) {
		resource[dependsOnFieldName] = resource.removeDependencyType(logger, vmasResourceType)
	}
	return resources
}

//RemoveNsgDependency Removes the nsg dependency from the resource
func RemoveNsgDependency(logger *logrus.Entry, resourceName string, resourceMap map[string]interface{}) {

	if resourceName != "" && resourceMap != nil {
		dependencies, ok := resourceMap[dependsOnFieldName].([]interface{})
		if !ok {
			logger.Warnf("Could not find dependencies for resourceName: %s", resourceName)
			return
		}

		for dIndex := len(dependencies) - 1; dIndex >= 0; dIndex-- {
			dependency := dependencies[dIndex].(string)
			if strings.Contains(dependency, nsgResourceType) || strings.Contains(dependency, nsgID) {
				dependencies = append(dependencies[:dIndex], dependencies[dIndex+1:]...)
			}
		}

		if len(dependencies) > 0 {
			resourceMap[dependsOnFieldName] = dependencies
		} else {
			delete(resourceMap, dependsOnFieldName)
		}

		return
	}
}

// NormalizeResourcesForK8sAgentUpgrade takes a template and removes elements that are unwanted in any scale/upgrade case
func (t *Transformer) NormalizeResourcesForK8sAgentUpgrade(logger *logrus.Entry, templateMap map[string]interface{}, isMasterManagedDisk bool, agentPoolsToPreserve map[string]bool) error {
	logger.Infoln("Running NormalizeResourcesForK8sMasterUpgrade....")
	if err := t.NormalizeResourcesForK8sMasterUpgrade(logger, templateMap, isMasterManagedDisk, agentPoolsToPreserve); err != nil {
		log.Fatalln(err)
		return err
	}

	logger.Infoln("Running NormalizeForK8sVMASScalingUp....")
	if err := t.NormalizeForK8sVMASScalingUp(logger, templateMap); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

// NormalizeForK8sAddVMASPool takes a template and removes elements that are unwanted in a K8s VMAS add pool case
func (t *Transformer) NormalizeForK8sAddVMASPool(l *logrus.Entry, templateMap map[string]interface{}) error {
	t.RemoveImmutableResourceProperties(l, templateMap)
	if err := t.RemoveJumpboxResourcesFromTemplate(l, templateMap); err != nil {
		return err
	}
	if err := t.NormalizeMasterResourcesForScaling(l, templateMap); err != nil {
		return err
	}
	if err := removeSingleOfType(l, templateMap, vnetResourceType); err != nil {
		return err
	}
	if err := removeSingleOfType(l, templateMap, rtResourceType); err != nil {
		return err
	}
	if err := removeSingleOfType(l, templateMap, nsgResourceType); err != nil {
		return err
	}
	return nil
}

// removeSingleOfType takes a template and removes references to resources of the given type
func removeSingleOfType(logger *logrus.Entry, templateMap map[string]interface{}, typeToRemove string) error {
	logger.Debugf("Looking for resources of type %s from the template.", typeToRemove)
	indexToRemove := -1

	templateResources := templateMap[resourcesFieldName].([]interface{})
	for i, r := range templateResources {
		resource, ok := r.(map[string]interface{})
		if !ok {
			logger.Warnf("Template improperly formatted for resource")
			continue
		}

		resourceType, found := resource[typeFieldName].(string)
		if found && resourceType == typeToRemove {
			if indexToRemove != -1 {
				err := errors.Errorf("Found at least 2 resources of type %s in the template but only 1 is expected", vnetResourceType)
				logger.Warnf(err.Error())
				return err
			}
			indexToRemove = i
		}

		deps, found := resource[dependsOnFieldName].([]interface{})
		if !found {
			continue
		}
		for idep := len(deps) - 1; idep >= 0; idep-- {
			dep := deps[idep].(string)
			if strings.Contains(dep, typeToRemove) || containsResourceID(dep, typeToRemove) {
				deps = append(deps[:idep], deps[idep+1:]...)
			}
		}
		if len(deps) > 0 {
			resource[dependsOnFieldName] = deps
		} else {
			delete(resource, dependsOnFieldName)
		}
	}

	if indexToRemove != -1 {
		logger.Debugf("Removing resource of type %s from the template.", typeToRemove)
		templateMap[resourcesFieldName] = removeIndexesFromArray(templateResources, []int{indexToRemove})
	} else {
		logger.Debugf("No resources of type %s were found.", typeToRemove)
	}
	return nil
}

func containsResourceID(dep, typeToRemove string) bool {
	switch typeToRemove {
	case nsgResourceType:
		return strings.Contains(dep, nsgID)
	case rtResourceType:
		return strings.Contains(dep, rtID)
	case vnetResourceType:
		return strings.Contains(dep, vnetID)
	default:
		return false
	}
}
