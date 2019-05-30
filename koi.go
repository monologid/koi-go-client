package koiclient

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
)

// IKoiClient represents available method for koi client
type IKoiClient interface {
	SetAPIUrl(url string) IKoiClient
	SetParams(param interface{}) IKoiClient
	Exec() ([]byte, error)
}

type client struct {
	APIKey string
	APIUrl string
	Param  interface{}
}

// New initiates koi client
func New(apiKey string) IKoiClient {
	return &client{
		APIKey: apiKey,
		APIUrl: "https://koi.monolog.id/api/v1/query",
	}
}

func (c *client) SetAPIUrl(url string) IKoiClient {
	c.APIUrl = url
	return c
}

func (c *client) SetParams(param interface{}) IKoiClient {
	c.Param = param
	return c
}

func (c *client) Exec() ([]byte, error) {
	var body []byte
	var errJSON error

	switch reflect.TypeOf(c.Param) {
	case reflect.TypeOf(&create{}):
		body, errJSON = c.Param.(ICreate).ToJSON()
	case reflect.TypeOf(&read{}):
		body, errJSON = c.Param.(IRead).ToJSON()
	case reflect.TypeOf(&update{}):
		body, errJSON = c.Param.(IUpdate).ToJSON()
	case reflect.TypeOf(&delete{}):
		body, errJSON = c.Param.(IDelete).ToJSON()
	default:
		errJSON = errors.New("invalid koi param type")
	}

	if errJSON != nil {
		return nil, errJSON
	}

	req, err := http.NewRequest(http.MethodPost, c.APIUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	authorizationToken := fmt.Sprintf("Bearer %s", c.APIKey)
	req.Header.Set("Authorization", authorizationToken)
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
