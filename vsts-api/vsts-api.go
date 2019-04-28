package vsts_api

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func ConstructAzureUri(organization string, project string, area string, resource string, version string) string {
	// {organization}/{project}/_apis/{area}/{resource}?api-version={version}
	return fmt.Sprintf(
		"https://dev.azure.com/%s/%s/_apis/%s/%s?api-version=%s", organization, project, area, resource, version,
	)
}

func ExecuteRequest(method string, uri string, personalAccessToken string, client *http.Client, requestBody io.Reader) *http.Response {
	req := constructRequest(method, uri, personalAccessToken, requestBody)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if (resp.StatusCode != 200) && (resp.StatusCode != 202) {
		log.Fatalf("API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	log.Printf("API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))

	return resp
}

func constructRequest(method string, uri string, personalAccessToken string, requestBody io.Reader) *http.Request {
	req, err := http.NewRequest(method, uri, nil)

	if method == "POST" {
		req, err = http.NewRequest(method, uri, requestBody)
	}

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("", personalAccessToken)
	return req
}
