package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func DecodeJson(response *http.Response, target interface{}) error {
	return json.NewDecoder(response.Body).Decode(target)
}
