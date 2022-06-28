package adapter

import "encoding/json"

func IsJson(data any) bool {
	_, err := json.Marshal(data)

	if err != nil {
		return false
	}
	
	return true
}