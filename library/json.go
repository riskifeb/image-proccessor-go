package library

import (
	"encoding/json"
	"fmt"
)

// Decode json data
func JsonDecode(content string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	error := json.Unmarshal([]byte(content), &result)
	if error != nil {
		return nil, fmt.Errorf("JSON Decode failed => %q", error)
	}
	return result, nil
}

func JsonEncode(content map[string]interface{}, indent bool) (string, error) {
	if indent {
		result, error := json.MarshalIndent(content, "", "\t")
		if error != nil {
			return "", fmt.Errorf("JSON Encode failed => %q", error)
		}
		return string(result), nil
	}
	result, error := json.Marshal(content)
	if error != nil {
		return "", fmt.Errorf("JSON Encode failed => %q", error)
	}
	return string(result), nil
}
