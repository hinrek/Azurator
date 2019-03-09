package project

import (
	"fmt"
	"github.com/hinrek/Azure-migrator/vsts-api"
)

type Project struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	URL      string `json:"url"`
	State    string `json:"state"`
	Revision int    `json:"revision"`
	Links    struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Collection struct {
			Href string `json:"href"`
		} `json:"collection"`
		Web struct {
			Href string `json:"href"`
		} `json:"web"`
	} `json:"_links"`
	Visibility  string `json:"visibility"`
	DefaultTeam struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"defaultTeam"`
	LastUpdateTime string `json:"lastUpdateTime"`
	Description    string `json:"description,omitempty"`
}

type Projects struct {
	Count   int       `json:"count"`
	Project []Project `json:"value"`
}

func List(organization string, apiVersion string) string {
	// https://dev.azure.com/{organization}/_apis/projects?api-version=5.0
	return vsts_api.ConstructAzureUrl(organization, "", "projects", apiVersion)
}

func Get(organization string, projectId string, apiVersion string) string {
	// https://dev.azure.com/{organization}/_apis/projects/{projectId}?api-version=5.0
	return vsts_api.ConstructAzureUrl(organization, "", fmt.Sprintf("projects/%s", projectId), apiVersion)
}
