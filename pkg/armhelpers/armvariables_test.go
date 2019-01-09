package armhelpers

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSizeMap(t *testing.T)  {
	sizeMap := getSizeMap()
	b, _ := json.MarshalIndent(sizeMap["vmSizesMap"], "", "   ")
	fmt.Println(string(b))
}
