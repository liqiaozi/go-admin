package utils

import "encoding/json"

func Object2JsonString(obj interface{}) string {
	result, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return string(result)
}
