package koiclient_test

import (
	"testing"

	koiclient "github.com/monologid/koi-go-client"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectSchemaForCreate(t *testing.T) {
	create := koiclient.NewCreate()
	create.SetSchema("sample_schema")

	assert.Equal(t, "sample_schema", create.GetSchema())
}

func TestShouldReturnCorrectValueForCreate(t *testing.T) {
	dummyValue := make(map[string]interface{})
	dummyValue["name"] = "ais"

	create := koiclient.NewCreate()
	create.SetValue(dummyValue)

	assert.Equal(t, dummyValue, create.GetValue())
}

func TestShouldReturnCorrectJSONSchemaForCreate(t *testing.T) {
	create := koiclient.NewCreate()
	create.SetSchema("sample_schema")

	dummyValue := make(map[string]interface{})
	dummyValue["name"] = "ais"
	create.SetValue(dummyValue)

	data, err := create.ToJSON()
	assert.NoError(t, err)
	assert.Equal(t, "{\"type\":\"CREATE\",\"schema\":\"sample_schema\",\"values\":{\"name\":\"ais\"}}", string(data))
}
