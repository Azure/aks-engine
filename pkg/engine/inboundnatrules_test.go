package engine

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateInboundNatRules(t *testing.T) {
	ibnr := createInboundNATRules()
	jsonObj, _ := json.MarshalIndent(ibnr, "", "   ")
	fmt.Println(string(jsonObj))
}
