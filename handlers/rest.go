package handlers

import (
	"encoding/json"
)

func SetRestMessage(req *Request) ([]byte, error) {
	message := map[string]string{
		"title":   req.Title,
		"link":    req.Link,
		"details": req.Details,
	}

	return json.Marshal(message)
}
