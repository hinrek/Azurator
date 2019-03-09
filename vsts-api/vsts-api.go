package vsts_api

import (
	"fmt"
	"log"
	"net/http"
)

func ConstructAzureUrl(teamProject string, area string, resource string, version string) string {
	return fmt.Sprintf(
		"https://dev.azure.com/%s/_apis/%s/%s?api-version=%s", teamProject, area, resource, version,
	)
}

func ResponseHandler(request string, personalAccessToken string, client *http.Client) *http.Response {
	req, err := requestHandler(request, personalAccessToken)
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		log.Fatalf("API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	} else if err != nil {
		log.Fatal(err)
	}
	return resp
}

func requestHandler(request string, personalAccessToken string) (*http.Request, error) {
	req, err := http.NewRequest("GET", request, nil)
	req.Header.Add("Content-Type", `"application/json"`)
	req.SetBasicAuth("", personalAccessToken)
	return req, err
}
