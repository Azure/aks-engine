// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package api

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/pkg/errors"
)

func checkJSONKeys(data []byte, types ...reflect.Type) error {
	var raw interface{}
	if e := json.Unmarshal(data, &raw); e != nil {
		return e
	}
	o := raw.(map[string]interface{})
	return checkMapKeys(o, types...)
}

func checkMapKeys(o map[string]interface{}, types ...reflect.Type) error {
	fieldMap := createJSONFieldMap(types)
	for k, v := range o {
		f, present := fieldMap[strings.ToLower(k)]
		if !present {
			return errors.Errorf("Unknown JSON tag %s", k)
		}
		if f.Type.Kind() == reflect.Struct && v != nil {
			if childMap, exists := v.(map[string]interface{}); exists {
				if e := checkMapKeys(childMap, f.Type); e != nil {
					return e
				}
			}
		}
		if f.Type.Kind() == reflect.Slice && v != nil {
			elementType := f.Type.Elem()
			if elementType.Kind() == reflect.Ptr {
				elementType = elementType.Elem()
			}
			if childSlice, exists := v.([]interface{}); exists {
				for _, child := range childSlice {
					if childMap, exists := child.(map[string]interface{}); exists {
						if e := checkMapKeys(childMap, elementType); e != nil {
							return e
						}
					}
				}
			}
		}
		if f.Type.Kind() == reflect.Ptr && v != nil {
			elementType := f.Type.Elem()
			if childMap, exists := v.(map[string]interface{}); exists {
				if e := checkMapKeys(childMap, elementType); e != nil {
					return e
				}
			}
		}
	}
	return nil
}

func createJSONFieldMap(types []reflect.Type) map[string]reflect.StructField {
	fieldMap := make(map[string]reflect.StructField)
	// Combine the permitted JSON fields from all types - handles the case
	// where we need to see the same JSON as a TypeMeta and a ContainerService
	for _, t := range types {
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			tag := f.Tag
			fieldJSON := tag.Get("json")
			if fieldJSON != "" {
				fieldJSONkey := strings.SplitN(fieldJSON, ",", 2)[0]
				fieldMap[strings.ToLower(fieldJSONkey)] = f
			}
		}
	}
	return fieldMap
}
