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

func (c *update) SetSchema(schema string) IUpdate {
	c.Schema = schema
	return c
}

func (c *update) GetSchema() string {
	return c.Schema
}

func (c *update) SetQuery(query map[string]interface{}) IUpdate {
	c.Query = query
	return c
}

func (c *update) GetQuery() map[string]interface{} {
	return c.Query
}

func (c *update) SetValue(value map[string]interface{}) IUpdate {
	c.Value = value
	return c
}

func (c *update) GetValue() map[string]interface{} {
	return c.Value
}

func (c *update) ToJSON() ([]byte, error) {
	jsonSchema := &updateSchema{
		Type:   "update",
		Schema: c.GetSchema(),
		Query:  c.GetQuery(),
		Values: c.GetValue(),
	}

	return json.Marshal(jsonSchema)
}
