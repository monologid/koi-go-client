package koiclient_test

import (
	"testing"

	koiclient "github.com/monologid/koi-go-client"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectSchemaForDelete(t *testing.T) {
	delete := koiclient.NewDelete()
	delete.SetSchema("sample_schema")

	assert.Equal(t, "sample_schema", delete.GetSchema())
}

func TestShouldReturnCorrectQueryForDelete(t *testing.T) {
	dummyQuery := make(map[string]interface{})
	dummyQuery["name"] = "ais"

	delete := koiclient.NewDelete()
	delete.SetQuery(dummyQuery)

	assert.Equal(t, dummyQuery, delete.GetQuery())
}

func TestShouldReturnCorrectJSONSchemaForDelete(t *testing.T) {
	delete := koiclient.NewDelete()
	delete.SetSchema("sample_schema")

	dummyQuery := make(map[string]interface{})
	dummyQuery["_id"] = "12345"
	delete.SetQuery(dummyQuery)

	data, err := delete.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, "{\"type\":\"DELETE\",\"schema\":\"sample_schema\",\"query\":{\"_id\":\"12345\"}}", string(data))
}
