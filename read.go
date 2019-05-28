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

func (c *read) SetSchema(schema string) IRead {
	c.Schema = schema
	return c
}

func (c *read) GetSchema() string {
	return c.Schema
}

func (c *read) SetAs(resultAs string) IRead {
	c.As = resultAs
	return c
}

func (c *read) GetAs() string {
	return c.As
}

func (c *read) SetField(fields []string) IRead {
	c.Field = fields
	return c
}

func (c *read) GetField() []string {
	return c.Field
}

func (c *read) SetQuery(query map[string]interface{}) IRead {
	c.Query = query
	return c
}

func (c *read) GetQuery() map[string]interface{} {
	return c.Query
}

func (c *read) SetLimit(limit int64) IRead {
	c.Limit = limit
	return c
}

func (c *read) GetLimit() int64 {
	return c.Limit
}

func (c *read) SetOffset(offset int64) IRead {
	c.Offset = offset
	return c
}

func (c *read) GetOffset() int64 {
	return c.Offset
}

func (c *read) SetOrderBy(orderBy []map[string]interface{}) IRead {
	c.OrderBy = orderBy
	return c
}

func (c *read) GetOrderBy() []map[string]interface{} {
	return c.OrderBy
}

func (c *read) ToJSON() ([]byte, error) {
	jsonSchema := &readSchema{
		Type:   "READ",
		Schema: c.GetSchema(),
		As:     c.GetAs(),
		Fields: c.GetField(),
		Query:  c.GetQuery(),
	}

	filter := &filters{}
	if c.GetLimit() > 0 {
		filter.Limit = c.GetLimit()
	} else {
		filter.Limit = 10
	}

	if c.GetOffset() > 0 {
		filter.Offset = c.GetOffset()
	}

	if len(c.GetOrderBy()) > 0 {
		filter.OrderBy = c.GetOrderBy()
	}

	jsonSchema.Filters = filter

	return json.Marshal(jsonSchema)
}
