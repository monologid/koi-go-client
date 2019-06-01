package koiclient

import (
	"encoding/json"
)

// IResult represents interface for result
type IResult interface {
	GetStatus() (string, error)
	GetData() (map[string]interface{}, error)
	GetDataAsList() ([]interface{}, error)
}

type result struct {
	Resp []byte
}

// NewResult initiates result object
func NewResult(resp []byte) IResult {
	return &result{
		Resp: resp,
	}
}

func (r *result) GetStatus() (string, error) {
	tempResult := make(map[string]interface{})

	err := json.Unmarshal(r.Resp, &tempResult)
	if err != nil {
		return "error", err
	}

	return tempResult["status"].(string), nil
}

func (r *result) GetData() (map[string]interface{}, error) {
	tempResult := make(map[string]interface{})

	err := json.Unmarshal(r.Resp, &tempResult)
	if err != nil {
		return nil, err
	}

	return tempResult["data"].(map[string]interface{}), nil
}

func (r *result) GetDataAsList() ([]interface{}, error) {
	tempResult := make(map[string]interface{})

	err := json.Unmarshal(r.Resp, &tempResult)
	if err != nil {
		return nil, err
	}

	return tempResult["data"].([]interface{}), nil
}
