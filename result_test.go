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
