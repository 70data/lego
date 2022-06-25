package structure

import (
	"encoding/json"
)

func CovertToJSON(element interface{}) string {
	tempJSON, _ := json.Marshal(element)
	return string(tempJSON)
}
