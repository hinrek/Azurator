package vsts_api

import (
	"fmt"
	"log"
	"net/http"
)

func ConstructAzureUri(organization string, project string, area string, resource string, version string) string {
	// {organization}/{project}/_apis/{area}/{resource}?api-version={version}
	return fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/%s/%s?api-version=%s", organization, project, area, resource, version,
	)
}

func ExecuteRequest(method string, uri string, personalAccessToken string, client *http.Client) *http.Response {
	req := constructRequest(method, uri, personalAccessToken)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return resp
}

func constructRequest(method string, uri string, personalAccessToken string) *http.Request {
	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Content-Type", `"application/json"`)
	req.SetBasicAuth("", personalAccessToken)

	return req
}
