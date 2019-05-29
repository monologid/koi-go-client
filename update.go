package koiclient

import "encoding/json"

// IUpdate represents update command for koi client
type IUpdate interface {
	SetSchema(schema string) IUpdate
	GetSchema() string
	SetQuery(query map[string]interface{}) IUpdate
	GetQuery() map[string]interface{}
	SetValue(value map[string]interface{}) IUpdate
	GetValue() map[string]interface{}
	ToJSON() ([]byte, error)
}

type update struct {
	Schema string
	Query  map[string]interface{}
	Value  map[string]interface{}
}

type updateSchema struct {
	Type   string                 `json:"type"`
	Schema string                 `json:"schema"`
	Query  map[string]interface{} `json:"query"`
	Values map[string]interface{} `json:"values"`
}

// NewUpdate initiates update feature
func NewUpdate() IUpdate {
	return &update{}
}

func (u *update) SetSchema(schema string) IUpdate {
	u.Schema = schema
	return u
}

func (u *update) GetSchema() string {
	return u.Schema
}

func (u *update) SetQuery(query map[string]interface{}) IUpdate {
	u.Query = query
	return u
}

func (u *update) GetQuery() map[string]interface{} {
	return u.Query
}

func (u *update) SetValue(value map[string]interface{}) IUpdate {
	u.Value = value
	return u
}

func (u *update) GetValue() map[string]interface{} {
	return u.Value
}

func (u *update) ToJSON() ([]byte, error) {
	jsonSchema := &updateSchema{
		Type:   "UPDATE",
		Schema: u.GetSchema(),
		Query:  u.GetQuery(),
		Values: u.GetValue(),
	}

	return json.Marshal(jsonSchema)
}
