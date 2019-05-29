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

func (d *delete) SetSchema(schema string) IDelete {
	d.Schema = schema
	return d
}

func (d *delete) GetSchema() string {
	return d.Schema
}

func (d *delete) SetQuery(query map[string]interface{}) IDelete {
	d.Query = query
	return d
}

func (d *delete) GetQuery() map[string]interface{} {
	return d.Query
}

func (d *delete) ToJSON() ([]byte, error) {
	jsonSchema := &deleteSchema{
		Type:   "DELETE",
		Schema: d.GetSchema(),
		Query:  d.GetQuery(),
	}

	return json.Marshal(jsonSchema)
}
