// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package engine

//import (
//	"encoding/json"
//	"fmt"
//	"testing"
//
//	"github.com/Azure/go-autorest/autorest/to"
//)
//
//func TestToArmParams(t *testing.T) {
//	aadTenantID := ARMParam{
//		DefaultValue: to.StringPtr(""),
//		Metadata: &ARMParamMetadata{
//			Description: to.StringPtr("The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription."),
//		},
//		Type: to.StringPtr("string"),
//	}
//
//	paramsMap := map[string]interface{}{}
//
//	paramsMap["aadTenantId"] = toARMParameterObj(aadTenantID)
//
//	jsonStr := map[string]interface{}{
//		"parameters": paramsMap,
//	}
//	jsonObj, _ := json.MarshalIndent(jsonStr, "", "   ")
//	fmt.Println(string(jsonObj))
//}
