package koiclient_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	koiclient "github.com/monologid/koi-go-client"
	"github.com/stretchr/testify/assert"
)

func TestShouldReceiveCreateRequest(t *testing.T) {
	query := koiclient.NewCreate()
	query.SetSchema("profile")

	value := make(map[string]interface{})
	value["name"] = "ais"
	value["email"] = "ais@mail.com"

	query.SetValue(value)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("body"))

		data, err := ioutil.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, "{\"type\":\"CREATE\",\"schema\":\"profile\",\"values\":{\"email\":\"ais@mail.com\",\"name\":\"ais\"}}", string(data))

		apiKey := req.Header.Get("Authorization")
		assert.Equal(t, "Bearer dummy-api-key", apiKey)
	}))
	defer func() { testServer.Close() }()

	client := koiclient.New("dummy-api-key")
	_, err := client.SetAPIUrl(testServer.URL).SetParams(query).Exec()

	assert.NoError(t, err)
}

func TestShouldReceiveReadRequest(t *testing.T) {
	query := koiclient.NewRead()
	query.SetSchema("profile")

	queryVal := make(map[string]interface{})
	queryVal["name"] = "ais"
	queryVal["email"] = "ais@mail.com"

	query.SetQuery(queryVal)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("body"))

		data, err := ioutil.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, "{\"type\":\"READ\",\"schema\":\"profile\",\"query\":{\"email\":\"ais@mail.com\",\"name\":\"ais\"},\"filters\":{\"limit\":10}}", string(data))

		apiKey := req.Header.Get("Authorization")
		assert.Equal(t, "Bearer dummy-api-key", apiKey)
	}))
	defer func() { testServer.Close() }()

	client := koiclient.New("dummy-api-key")
	_, err := client.SetAPIUrl(testServer.URL).SetParams(query).Exec()

	assert.NoError(t, err)
}

func TestShouldReceiveUpdateRequest(t *testing.T) {
	query := koiclient.NewUpdate()
	query.SetSchema("profile")
	query.SetQuery(map[string]interface{}{"_id": "12345"})

	value := make(map[string]interface{})
	value["name"] = "ais"
	value["email"] = "ais@mail.com"
	query.SetValue(value)

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("body"))

		data, err := ioutil.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, "{\"type\":\"UPDATE\",\"schema\":\"profile\",\"query\":{\"_id\":\"12345\"},\"values\":{\"email\":\"ais@mail.com\",\"name\":\"ais\"}}", string(data))

		apiKey := req.Header.Get("Authorization")
		assert.Equal(t, "Bearer dummy-api-key", apiKey)
	}))
	defer func() { testServer.Close() }()

	client := koiclient.New("dummy-api-key")
	_, err := client.SetAPIUrl(testServer.URL).SetParams(query).Exec()

	assert.NoError(t, err)
}

func TestShouldReceiveDeleteRequest(t *testing.T) {
	query := koiclient.NewDelete()
	query.SetSchema("profile")
	query.SetQuery(map[string]interface{}{"_id": "12345"})

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("body"))

		data, err := ioutil.ReadAll(req.Body)
		assert.NoError(t, err)
		assert.Equal(t, "{\"type\":\"DELETE\",\"schema\":\"profile\",\"query\":{\"_id\":\"12345\"}}", string(data))

		apiKey := req.Header.Get("Authorization")
		assert.Equal(t, "Bearer dummy-api-key", apiKey)
	}))
	defer func() { testServer.Close() }()

	client := koiclient.New("dummy-api-key")
	_, err := client.SetAPIUrl(testServer.URL).SetParams(query).Exec()

	assert.NoError(t, err)
}

func TestShouldReturnErrorIfTypeIsInvalid(t *testing.T) {
	query := "test"
	client := koiclient.New("dummy-api-key")
	_, err := client.SetParams(query).Exec()

	assert.Error(t, err)
}

func TestShouldReturnErrorIfAPIURLIsNotValid(t *testing.T) {
	query := koiclient.NewDelete()
	query.SetSchema("profile")
	query.SetQuery(map[string]interface{}{"_id": "12345"})

	client := koiclient.New("dummy-api-key")
	_, err := client.SetAPIUrl("-").SetParams(query).Exec()

	assert.Error(t, err)
}
