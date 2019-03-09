package vsts_api

import (
	"fmt"
	"log"
	"net/http"
)

func ConstructAzureUrl(teamProject string, area string, resource string, version string) string {

	url := fmt.Sprintf("https://dev.azure.com/%s/_apis/%s/%s?api-version=%s", teamProject, area, resource, version)

	return url
}

func RequestHandler(project, personalAccessToken string, client *http.Client) *http.Response {
	//TODO split the request and response into different methods
	req, err := http.NewRequest("GET", fmt.Sprintf(project), nil)
	req.Header.Add("Content-Type", `"application/json"`)
	req.SetBasicAuth("", personalAccessToken)
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		log.Fatalf("API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	} else if err != nil {
		log.Fatal(err)
	}
	return resp
}
