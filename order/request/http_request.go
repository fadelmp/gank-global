package request

import (
	"io"
	"io/ioutil"
	"net/http"
)

type HttpRequest interface {
	GetRequest(url string) (string, error)
	PostRequest(url string, body []byte) (string, error)
}

func GetRequest(url string) (result string, err error) {

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func PostRequest(url string, param_body io.Reader) (result string, err error) {

	app_type := "application/json"

	resp, err := http.Post(url, app_type, param_body)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
