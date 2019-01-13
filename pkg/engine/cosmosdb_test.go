package engine

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCreateCosmosDB(t *testing.T) {
	db := createCosmosDBAccount()

	jsonObj, _ := json.MarshalIndent(db, "", "   ")
	fmt.Println(string(jsonObj))
}

