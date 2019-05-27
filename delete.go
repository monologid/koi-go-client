package koiclient

import "encoding/json"

// IDelete represents delete command for koi client
type IDelete interface {
	SetSchema(schema string) IDelete
	GetSchema() string
	SetQuery(query map[string]interface{}) IDelete
	GetQuery() map[string]interface{}
	ToJSON() ([]byte, error)
}

type delete struct {
	Schema string
	Query  map[string]interface{}
}

type deleteSchema struct {
	Type   string                 `json:"type"`
	Schema string                 `json:"schema"`
	Query  map[string]interface{} `json:"query"`
}

// NewDelete initiates delete feature
func NewDelete() IDelete {
	return &delete{}
}

func (c *delete) SetSchema(schema string) IDelete {
	c.Schema = schema
	return c
}

func (c *delete) GetSchema() string {
	return c.Schema
}

func (c *delete) SetQuery(query map[string]interface{}) IDelete {
	c.Query = query
	return c
}

func (c *delete) GetQuery() map[string]interface{} {
	return c.Query
}

func (c *delete) ToJSON() ([]byte, error) {
	jsonSchema := &deleteSchema{
		Type:   "delete",
		Schema: c.GetSchema(),
		Query:  c.GetQuery(),
	}

	return json.Marshal(jsonSchema)
}
