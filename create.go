package koiclient

import "encoding/json"

// ICreate represents create command for koi client
type ICreate interface {
	SetSchema(schema string) ICreate
	GetSchema() string
	SetValue(value map[string]interface{}) ICreate
	GetValue() map[string]interface{}
	ToJSON() ([]byte, error)
}

type create struct {
	Schema string
	Value  map[string]interface{}
}

type createSchema struct {
	Type   string                 `json:"type"`
	Schema string                 `json:"schema"`
	Values map[string]interface{} `json:"values"`
}

// NewCreate initiates create feature
func NewCreate() ICreate {
	return &create{}
}

func (c *create) SetSchema(schema string) ICreate {
	c.Schema = schema
	return c
}

func (c *create) GetSchema() string {
	return c.Schema
}

func (c *create) SetValue(value map[string]interface{}) ICreate {
	c.Value = value
	return c
}

func (c *create) GetValue() map[string]interface{} {
	return c.Value
}

func (c *create) ToJSON() ([]byte, error) {
	jsonSchema := &createSchema{
		Type:   "CREATE",
		Schema: c.GetSchema(),
		Values: c.GetValue(),
	}

	return json.Marshal(jsonSchema)
}
