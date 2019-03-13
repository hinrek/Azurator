package project

import (
	"github.com/hinrek/Azure-migrator/utils"
	"github.com/hinrek/Azure-migrator/vsts-api"
	"log"
	"net/http"
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

func (projects *Projects) List(organization string, apiVersion string, personalAccessToken string, client *http.Client) *Projects {
	// https://dev.azure.com/{organization}/_apis/projects?api-version=5.0
	url := vsts_api.ConstructAzureUrl(organization, "", "projects", "", apiVersion)
	httpResponse := vsts_api.ResponseHandler(url, personalAccessToken, client)

	bytes, err := utils.ReadResponseBody(*httpResponse)
	if err != nil {
		log.Fatal(err)
	}
	err = utils.JsonUnmarshal(bytes, &projects)
	if err != nil {
		log.Fatal(err)
	}

	return projects
}

func (project *Project) Get(organization string, projectId string, apiVersion string, personalAccessToken string, client *http.Client) *Project {
	// https://dev.azure.com/{organization}/_apis/projects/{projectId}?api-version=5.0
	url := vsts_api.ConstructAzureUrl(organization, "", "projects", projectId, apiVersion)
	httpResponse := vsts_api.ResponseHandler(url, personalAccessToken, client)

	bytes, err := utils.ReadResponseBody(*httpResponse)
	if err != nil {
		log.Fatal(err)
	}
	err = utils.JsonUnmarshal(bytes, &project)
	if err != nil {
		log.Fatal(err)
	}

	return project
}
