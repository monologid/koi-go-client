package koiclient_test

import (
	"testing"

	koiclient "github.com/monologid/koi-go-client"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnCorrectStatus(t *testing.T) {
	response := `{"status": "success", "data": {}}`
	result := koiclient.NewResult([]byte(response))

	status, err := result.GetStatus()

	assert.NoError(t, err)
	assert.Equal(t, "success", status)
}

func TestShouldReturnErrorWhenJSONIsNotValidForStatus(t *testing.T) {
	response := `{"status": "success", "data": {}`
	result := koiclient.NewResult([]byte(response))

	status, err := result.GetStatus()

	assert.Error(t, err)
	assert.Equal(t, "error", status)
}

func TestShouldReturnCorrectData(t *testing.T) {
	response := `{"status": "success", "data": {"name":"monolog"}}`
	result := koiclient.NewResult([]byte(response))

	status, err := result.GetStatus()
	assert.NoError(t, err)
	assert.Equal(t, "success", status)

	data, err := result.GetData()
	assert.NoError(t, err)
	assert.Equal(t, "monolog", data["name"])
}

func TestShouldReturnErrorWhenJSONIsNotValidForData(t *testing.T) {
	response := `{"status": "success", "data": {"name":"monolog"}`
	result := koiclient.NewResult([]byte(response))

	_, err := result.GetData()
	assert.Error(t, err)
}

func TestShouldReturnCorrectDataAsList(t *testing.T) {
	response := `{"status":"success","data":[{"_id":"2c3e6b70-8427-11e9-ac2c-a9ad4a2ad58d","account_id":"12345","app_secret_key":"7b50b4a9-08a4-4fea-8ab4-92a3da76186f","description":"test descruiption","endpoint":"/api/v1/sample","filename":"test.v2.so","method":"GET","module_name":"test","name":"test"}]}`
	result := koiclient.NewResult([]byte(response))

	data, err := result.GetDataAsList()
	assert.NoError(t, err)
	assert.Equal(t, 1, len(data))
}

func TestShouldReturnErrorOnGetDataAsList(t *testing.T) {
	response := `{"status":"success","data":[{"_id":"2c3e6b70-8427-11e9-ac2c-a9ad4a2ad58d","account_id":"12345","app_secret_key":"7b50b4a9-08a4-4fea-8ab4-92a3da76186f","description":"test descruiption","endpoint":"/api/v1/sample","filename":"test.v2.so","method":"GET","module_name":"test","name":"test"}}`
	result := koiclient.NewResult([]byte(response))

	_, err := result.GetDataAsList()
	assert.Error(t, err)
}
