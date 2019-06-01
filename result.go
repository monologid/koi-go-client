package koiclient

import (
	"encoding/json"
)

// IResult represents interface for result
type IResult interface {
	Status() string
	Data() interface{}
	DataAsList() []interface{}
}

type result struct {
	Body map[string]interface{}
}

// NewResult initiates result object
func NewResult(resp []byte) (IResult, error) {
	tempResult := make(map[string]interface{})
	err := json.Unmarshal(resp, &tempResult)
	if err != nil {
		return nil, err
	}

	return &result{
		Body: tempResult,
	}, nil
}

func (r *result) Status() string {
	return r.Body["status"].(string)
}

func (r *result) Data() interface{} {
	return r.Body["data"].(interface{})
}

func (r *result) DataAsList() []interface{} {
	res, ok := r.Body["data"].([]interface{})
	if !ok {
		return nil
	}
	return res
}
