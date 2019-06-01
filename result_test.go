package koiclient_test

import (
	"testing"

	koiclient "github.com/monologid/koi-go-client"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnErrorIfResponseIsBroken(t *testing.T) {
	response := `{"status": "success", "data": {}`
	_, err := koiclient.NewResult([]byte(response))

	assert.Error(t, err)
}

func TestShouldReturnCorrectStatus(t *testing.T) {
	response := `{"status": "success", "data": {}}`
	result, err := koiclient.NewResult([]byte(response))

	assert.NoError(t, err)
	assert.Equal(t, "success", result.Status())
}

func TestShouldReturnCorrectEmptyData(t *testing.T) {
	response := `{"status": "success", "data": {}}`
	result, err := koiclient.NewResult([]byte(response))

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{}, result.Data())
}

func TestShouldReturnCorrectData(t *testing.T) {
	response := `{"status": "success", "data": {"name":"monolog"}}`
	result, err := koiclient.NewResult([]byte(response))

	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"name": "monolog"}, result.Data())
}

func TestShouldReturnCorrectDataAsList(t *testing.T) {
	response := `{"status":"success","data":[{"_id":"2c3e6b70-8427-11e9-ac2c-a9ad4a2ad58d","account_id":"12345","app_secret_key":"7b50b4a9-08a4-4fea-8ab4-92a3da76186f","description":"test descruiption","endpoint":"/api/v1/sample","filename":"test.v2.so","method":"GET","module_name":"test","name":"test"}]}`
	result, err := koiclient.NewResult([]byte(response))

	assert.NoError(t, err)
	assert.Equal(t, 1, len(result.DataAsList()))
}

func TestShouldReturnZeroIfDataIsNotListButUseDataAsList(t *testing.T) {
	response := `{"status":"success","data":{}}`
	result, err := koiclient.NewResult([]byte(response))

	assert.NoError(t, err)
	assert.Equal(t, 0, len(result.DataAsList()))
}
