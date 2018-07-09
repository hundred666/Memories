package model

import "encoding/json"

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

func MarshalResponse(code int, payload interface{}) ([]byte) {
	resp := Response{Status: code, Payload: payload}
	result, err := json.Marshal(resp)
	if err != nil {
		return nil
	}
	return result
}
