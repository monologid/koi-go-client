package koiclient_test

import (
	"testing"

	koiclient "github.com/monologid/koi-go-client"
	"github.com/stretchr/testify/assert"
)

var data = `{
  "type": "READ",
  "schema": "account_profile",
  "as": "account",
  "fields": [
    "email",
    "phone_no",
    "description"
  ],
  "query": {
    "email": "john@doe.com"
  },
  "filters": {
    "limit": 20,
    "offset": 1,
    "order_by": [
      { "email": "ASC" },
      { "description": "DESC" }
    ]
  }
}`

func TestShouldReturnCorrectSchemaForRead(t *testing.T) {
	read := koiclient.NewRead()
	read.SetSchema("sample_schema")

	assert.Equal(t, "sample_schema", read.GetSchema())
}

func TestShouldReturnCorrectResultAsForRead(t *testing.T) {
	read := koiclient.NewRead()
	read.SetAs("sample_schema")

	assert.Equal(t, "sample_schema", read.GetAs())
}

func TestShouldReturnCorrectFieldForRead(t *testing.T) {
	read := koiclient.NewRead()
	read.SetField([]string{"name", "age"})

	assert.Equal(t, []string{"name", "age"}, read.GetField())
}

func TestShouldReturnCorrectQueryForRead(t *testing.T) {
	dummyQuery := make(map[string]interface{})
	dummyQuery["name"] = "ais"

	read := koiclient.NewRead()
	read.SetQuery(dummyQuery)

	assert.Equal(t, dummyQuery, read.GetQuery())
}

func TestShouldReturnCorrectLimitForRead(t *testing.T) {
	read := koiclient.NewRead()
	read.SetLimit(10)

	assert.Equal(t, int64(10), read.GetLimit())
}

func TestShouldReturnCorrectOffsetForRead(t *testing.T) {
	read := koiclient.NewRead()
	read.SetOffset(5)

	assert.Equal(t, int64(5), read.GetOffset())
}

func TestShouldReturnCorrectOrderByForRead(t *testing.T) {
	read := koiclient.NewRead()

	var orderBy []map[string]interface{}
	orderBy = append(orderBy, map[string]interface{}{"email": "ASC"})

	read.SetOrderBy(orderBy)

	assert.Equal(t, orderBy, read.GetOrderBy())
}

func TestShouldReturnCorrectJSONSchemaForRead(t *testing.T) {
	read := koiclient.NewRead()
	read.SetSchema("sample_schema")
	read.SetAs("sample_schema")

	fields := []string{"name", "age"}
	read.SetField(fields)

	dummyQuery := make(map[string]interface{})
	dummyQuery["_id"] = "12345"
	read.SetQuery(dummyQuery)

	data, err := read.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, "{\"type\":\"READ\",\"schema\":\"sample_schema\",\"as\":\"sample_schema\",\"fields\":[\"name\",\"age\"],\"query\":{\"_id\":\"12345\"},\"filters\":{}}", string(data))
}
