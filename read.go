package koiclient

import "encoding/json"

// IRead represents read command for koi client
type IRead interface {
	SetSchema(schema string) IRead
	GetSchema() string
	SetAs(resultAs string) IRead
	GetAs() string
	SetField(fields []string) IRead
	GetField() []string
	SetQuery(query map[string]interface{}) IRead
	GetQuery() map[string]interface{}
	SetLimit(limit int64) IRead
	GetLimit() int64
	SetOffset(offset int64) IRead
	GetOffset() int64
	SetOrderBy(orderBy []map[string]interface{}) IRead
	GetOrderBy() []map[string]interface{}
	ToJSON() ([]byte, error)
}

type read struct {
	Schema  string
	As      string
	Field   []string
	Query   map[string]interface{}
	Value   map[string]interface{}
	Limit   int64
	Offset  int64
	OrderBy []map[string]interface{}
}

type readSchema struct {
	Type    string                 `json:"type"`
	Schema  string                 `json:"schema"`
	As      string                 `json:"as,omitempty"`
	Fields  []string               `json:"fields,omitempty"`
	Query   map[string]interface{} `json:"query"`
	Filters *filters               `json:"filters,omitempty"`
}

type filters struct {
	Limit   int64                    `json:"limit,omitempty"`
	Offset  int64                    `json:"offset,omitempty"`
	OrderBy []map[string]interface{} `json:"order_by,omitempty"`
}

// NewRead initiates read feature
func NewRead() IRead {
	return &read{}
}

func (r *read) SetSchema(schema string) IRead {
	r.Schema = schema
	return r
}

func (r *read) GetSchema() string {
	return r.Schema
}

func (r *read) SetAs(resultAs string) IRead {
	r.As = resultAs
	return r
}

func (r *read) GetAs() string {
	return r.As
}

func (r *read) SetField(fields []string) IRead {
	r.Field = fields
	return r
}

func (r *read) GetField() []string {
	return r.Field
}

func (r *read) SetQuery(query map[string]interface{}) IRead {
	r.Query = query
	return r
}

func (r *read) GetQuery() map[string]interface{} {
	return r.Query
}

func (r *read) SetLimit(limit int64) IRead {
	r.Limit = limit
	return r
}

func (r *read) GetLimit() int64 {
	return r.Limit
}

func (r *read) SetOffset(offset int64) IRead {
	r.Offset = offset
	return r
}

func (r *read) GetOffset() int64 {
	return r.Offset
}

func (r *read) SetOrderBy(orderBy []map[string]interface{}) IRead {
	r.OrderBy = orderBy
	return r
}

func (r *read) GetOrderBy() []map[string]interface{} {
	return r.OrderBy
}

func (r *read) ToJSON() ([]byte, error) {
	jsonSchema := &readSchema{
		Type:   "READ",
		Schema: r.GetSchema(),
		As:     r.GetAs(),
		Fields: r.GetField(),
		Query:  r.GetQuery(),
	}

	filter := &filters{}
	if r.GetLimit() > 0 {
		filter.Limit = r.GetLimit()
	} else {
		filter.Limit = 10
	}

	if r.GetOffset() > 0 {
		filter.Offset = r.GetOffset()
	}

	if len(r.GetOrderBy()) > 0 {
		filter.OrderBy = r.GetOrderBy()
	}

	jsonSchema.Filters = filter

	return json.Marshal(jsonSchema)
}
