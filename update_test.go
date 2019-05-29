package koiclient_test

import (
	"testing"

	koiclient "github.com/monologid/koi-go-client"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectSchemaForUpdate(t *testing.T) {
	update := koiclient.NewUpdate()
	update.SetSchema("sample_schema")

	assert.Equal(t, "sample_schema", update.GetSchema())
}

func TestShouldReturnCorrectValueForUpdate(t *testing.T) {
	dummyValue := make(map[string]interface{})
	dummyValue["name"] = "ais"

	update := koiclient.NewUpdate()
	update.SetValue(dummyValue)

	assert.Equal(t, dummyValue, update.GetValue())
}

func TestShouldReturnCorrectQueryForUpdate(t *testing.T) {
	dummyQuery := make(map[string]interface{})
	dummyQuery["name"] = "ais"

	update := koiclient.NewUpdate()
	update.SetQuery(dummyQuery)

	assert.Equal(t, dummyQuery, update.GetQuery())
}

func TestShouldReturnCorrectJSONSchemaForUpdate(t *testing.T) {
	update := koiclient.NewUpdate()
	update.SetSchema("sample_schema")

	dummyQuery := make(map[string]interface{})
	dummyQuery["_id"] = "12345"
	update.SetQuery(dummyQuery)

	dummyValue := make(map[string]interface{})
	dummyValue["name"] = "ais"
	update.SetValue(dummyValue)

	data, err := update.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, "{\"type\":\"UPDATE\",\"schema\":\"sample_schema\",\"query\":{\"_id\":\"12345\"},\"values\":{\"name\":\"ais\"}}", string(data))
}
