package transform

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"

	"github.com/Jeffail/gabs"
	log "github.com/sirupsen/logrus"
)

// APIModelValue represents a value in the APIModel JSON file
type APIModelValue struct {
	stringValue   string
	intValue      int64
	arrayValue    bool
	arrayIndex    int
	arrayProperty string
	arrayName     string
}

// MapValues converts an arraw of rwa ApiModel values (like ["masterProfile.count=4","linuxProfile.adminUsername=admin"]) to a map
func MapValues(m map[string]APIModelValue, setFlagValues []string) {
	if len(setFlagValues) == 0 {
		return
	}

	// regex to find array[index].property pattern in the key, like linuxProfile.ssh.publicKeys[0].keyData
	re := regexp.MustCompile(`(.*?)\[(.*?)\]\.(.*?)$`)

	for _, setFlagValue := range setFlagValues {
		kvpMap := parseKeyValuePairs(setFlagValue)
		for key, keyValue := range kvpMap {
			flagValue := APIModelValue{}
			// try to parse the value as integer or fallback to string
			if keyValueAsInteger, err := strconv.ParseInt(keyValue, 10, 64); err == nil {
				flagValue.intValue = keyValueAsInteger
			} else {
				flagValue.stringValue = keyValue
			}

			// check if the key is an array property
			keyArrayMatch := re.FindStringSubmatch(key)

			// it's an array
			if keyArrayMatch != nil {
				i, err := strconv.ParseInt(keyArrayMatch[2], 10, 32)
				if err != nil {
					log.Warnln(fmt.Sprintf("array index is not specified for property %s", key))
				} else {
					arrayIndex := int(i)
					flagValue.arrayValue = true
					flagValue.arrayName = keyArrayMatch[1]
					flagValue.arrayIndex = arrayIndex
					flagValue.arrayProperty = keyArrayMatch[3]
					m[key] = flagValue
				}
			} else {
				m[key] = flagValue
			}
		}
	}
}

// MergeValuesWithAPIModel takes the path to an ApiModel JSON file, loads it and merges it with the values in the map to another temp file
func MergeValuesWithAPIModel(apiModelPath string, m map[string]APIModelValue) (string, error) {
	// load the apiModel file from path
	fileContent, err := ioutil.ReadFile(apiModelPath)
	if err != nil {
		return "", err
	}

	// parse the json from file content
	jsonObj, err := gabs.ParseJSON(fileContent)
	if err != nil {
		return "", err
	}

	// update api model definition with each value in the map
	for key, flagValue := range m {
		// working on an array
		if flagValue.arrayValue {
			log.Debugln(fmt.Sprintf("--set flag array value detected. Name: %s, Index: %b, PropertyName: %s", flagValue.arrayName, flagValue.arrayIndex, flagValue.arrayProperty))
			arrayValue := jsonObj.Path(fmt.Sprint("properties.", flagValue.arrayName))
			if flagValue.stringValue != "" {
				arrayValue.Index(flagValue.arrayIndex).SetP(flagValue.stringValue, flagValue.arrayProperty)
			} else {
				arrayValue.Index(flagValue.arrayIndex).SetP(flagValue.intValue, flagValue.arrayProperty)
			}
		} else {
			if flagValue.stringValue != "" {
				jsonObj.SetP(flagValue.stringValue, fmt.Sprint("properties.", key))
			} else {
				jsonObj.SetP(flagValue.intValue, fmt.Sprint("properties.", key))
			}
		}
	}

	// generate a new file
	tmpFile, err := ioutil.TempFile("", "mergedApiModel")
	if err != nil {
		return "", err
	}

	tmpFileName := tmpFile.Name()
	err = ioutil.WriteFile(tmpFileName, []byte(jsonObj.String()), os.ModeAppend)
	if err != nil {
		return "", err
	}

	return tmpFileName, nil
}

func parseKeyValuePairs(literal string) map[string]string {
	log.Debugln(fmt.Sprintf("parsing --set flag key/value pairs from %s", literal))
	inQuoteLiteral := false
	inDblQuoteLiteral := false
	inKey := true
	kvpMap := map[string]string{}

	currentKey := ""
	currentValue := ""

	for _, literalChar := range literal {
		switch literalChar {
		case '\'': // if we hit a ' char
			if !inQuoteLiteral && !inDblQuoteLiteral { // and we are not already in a literal
				inQuoteLiteral = true // start a new ' delimited literal value
				inKey = false
			} else if inQuoteLiteral { // we already are in a ' delimited literal value
				inQuoteLiteral = false // stop it
				inKey = true
			}
		case '"': // if we hit a " char
			if !inDblQuoteLiteral && !inQuoteLiteral { // and we are not already in a literal
				inDblQuoteLiteral = true // start a new " delimited literal value
				inKey = false
			} else if inDblQuoteLiteral { // we already are in a " delimited literal value
				inDblQuoteLiteral = false // stop it
				inKey = true
			}
		case ',': // if we hit a , char
			if inQuoteLiteral || inDblQuoteLiteral { // we are in a literal
				currentValue += string(literalChar)
			} else {
				log.Debugln(fmt.Sprintf("new key/value parsed: %s = %s", currentKey, currentValue))
				kvpMap[currentKey] = currentValue
				currentKey = ""
				currentValue = ""
				inKey = true
			}
		case '=': // if we hit a = char
			if inQuoteLiteral || inDblQuoteLiteral || !inKey { // we are in a literal / value
				currentValue += string(literalChar)
			} else {
				inKey = false
			}
		default: // we hit any other char
			if inKey {
				currentKey += string(literalChar)
			} else {
				currentValue += string(literalChar)
			}
		}
	}

	// push latest literal
	if currentKey != "" {
		log.Debugln(fmt.Sprintf("new key/value parsed: %s = %s", currentKey, currentValue))
		kvpMap[currentKey] = currentValue
	}

	return kvpMap
}
