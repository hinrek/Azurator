package project

import (
	"bytes"
	"encoding/json"
	"github.com/hinrek/Azure-migrator/utils"
	"github.com/hinrek/Azure-migrator/vsts-api"
	"log"
	"net/http"
)

type Projects struct {
	Count   int       `json:"count"`
	Project []Project `json:"value"`
}

type Project struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	URL            string       `json:"url"`
	State          string       `json:"state"`
	Capabilities   Capabilities `json:"capabilities"`
	Revision       int          `json:"revision"`
	Links          Links        `json:"_links"`
	Visibility     string       `json:"visibility"`
	DefaultTeam    DefaultTeam  `json:"defaultTeam"`
	LastUpdateTime string       `json:"lastUpdateTime"`
	Description    string       `json:"description,omitempty"`
}

type ProcessTemplate struct {
	TemplateName   string `json:"templateName"`
	TemplateTypeID string `json:"templateTypeId"`
}

type Versioncontrol struct {
	SourceControlType string `json:"sourceControlType"`
	GitEnabled        string `json:"gitEnabled"`
	TfvcEnabled       string `json:"tfvcEnabled"`
}

type Capabilities struct {
	ProcessTemplate ProcessTemplate `json:"processTemplate"`
	Versioncontrol  Versioncontrol  `json:"versioncontrol"`
}

type DefaultTeam struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Links struct {
	Self       Self       `json:"self"`
	Collection Collection `json:"collection"`
	Web        Web        `json:"web"`
}

type Self struct {
	Href string `json:"href"`
}

type Collection struct {
	Href string `json:"href"`
}

type Web struct {
	Href string `json:"href"`
}

func (projects *Projects) List(organization string, apiVersion string, personalAccessToken string, client *http.Client) *Projects {
	// https://dev.azure.com/{organization}/_apis/projects?api-version=5.0
	url := vsts_api.ConstructAzureUri(organization, "", "projects", "", apiVersion)
	httpResponse := vsts_api.ExecuteRequest("GET", url, personalAccessToken, client, nil)

	err := utils.DecodeJson(httpResponse, projects)
	if err != nil {
		log.Fatal(err)
	}

	return projects
}

func (project *Project) Get(organization string, projectId string, apiVersion string, personalAccessToken string, client *http.Client) *Project {
	// GET https://dev.azure.com/{organization}/_apis/projects/{projectId}?includeCapabilities={includeCapabilities}&api-version=5.0
	url := vsts_api.ConstructAzureUri(organization, "", "projects", projectId+"?includeCapabilities=true&includeHistory=true", apiVersion)
	httpResponse := vsts_api.ExecuteRequest("GET", url, personalAccessToken, client, nil)

	err := utils.DecodeJson(httpResponse, project)
	if err != nil {
		log.Fatal(err)
	}

	return project
}

func (project *Project) Create(organization string, apiVersion string, personalAccessToken string, client *http.Client) *http.Response {
	// POST https://dev.azure.com/{organization}/_apis/projects?api-version=5.0
	url := vsts_api.ConstructAzureUri(organization, "", "projects", "", apiVersion)

	var postProject = struct {
		Name         string `json:"name"`
		Description  string `json:"description"`
		Capabilities struct {
			ProcessTemplate struct {
				TemplateTypeID string `json:"templateTypeId"`
			} `json:"processTemplate"`
			Versioncontrol struct {
				SourceControlType string `json:"sourceControlType"`
			} `json:"versioncontrol"`
		} `json:"capabilities"`
		Visibility string `json:"visibility"`
	}{}

	postProject.Name = project.Name
	postProject.Description = project.Description
	postProject.Capabilities.Versioncontrol.SourceControlType = project.Capabilities.Versioncontrol.SourceControlType
	postProject.Capabilities.ProcessTemplate.TemplateTypeID = project.Capabilities.ProcessTemplate.TemplateTypeID
	postProject.Visibility = project.Visibility

	body, err := json.Marshal(postProject)
	if err != nil {
		log.Println(err)
	}

	httpResponse := vsts_api.ExecuteRequest("POST", url, personalAccessToken, client, bytes.NewBuffer(body))

	return httpResponse
}
