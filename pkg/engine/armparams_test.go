package engine

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/go-autorest/autorest/to"
	"testing"
)

func TestToArmParams(t *testing.T) {
	aadTenantId := ARMParam{
		DefaultValue: to.StringPtr(""),
		Metadata: &ARMParamMetadata{
			Description: to.StringPtr("The AAD tenant ID to use for authentication. If not specified, will use the tenant of the deployment subscription."),
		},
		Type: to.StringPtr("string"),
	}

	paramsMap := map[string]interface{}{}

	paramsMap["aadTenantId"] = toARMParameterObj(aadTenantId)

	jsonStr := map[string]interface{}{
		"parameters": paramsMap,
	}
	jsonObj, _ := json.MarshalIndent(jsonStr, "", "   ")
	fmt.Println(string(jsonObj))
}
