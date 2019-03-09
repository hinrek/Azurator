package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func ReadResponseBody(response http.Response) ([]byte, error) {
	return ioutil.ReadAll(response.Body)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
