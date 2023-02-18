package Config

import "encoding/json"

func IsJson(message []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(message, &js) == nil
}
