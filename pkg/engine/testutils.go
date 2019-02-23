// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

import "github.com/fatih/structs"

func resourceSliceToMap(resources []interface{}) map[string]interface{} {
	resourceMap := map[string]interface{}{}
	for _, resource := range resources {
		if isARMResource(resource) {
			resourceName := extractNameFromARMType(resource)
			resourceMap[resourceName] = resource
		}
	}
	return resourceMap
}

func extractNameFromARMType(resource interface{}) string {
	s := structs.New(resource)
	fields := s.Field(s.Names()[1]).Fields()
	for _, f := range fields {
		innerField := s.Field(f.Name())
		innerFieldValue := innerField.Value()
		if innerStringValue, ok := innerFieldValue.(*string); ok && innerField.Name() == "Name" {
			return *innerStringValue
		}
	}
	return ""
}

func isARMResource(resource interface{}) bool {
	s := structs.New(resource)
	_, ok := s.FieldOk("ARMResource")
	return ok
}
