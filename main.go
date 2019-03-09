package main

import (
	"encoding/json"
	"fmt"
	"github.com/hinrek/Azure-migrator/vsts-api"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hinrek/Azure-migrator/models/configuration"
	"github.com/hinrek/Azure-migrator/models/project"
)

var (
	conf          configuration.Conf
	projects      project.Projects
	singleProject project.Project
)

func init() {
	configFilePath := "configs/config.yml"
	conf.Get(configFilePath)
}

func main() {
	organization := conf.SourceOrganization.Name
	apiVersion := conf.SourceOrganization.APIVersion
	personalAccessToken := conf.SourceOrganization.PersonalAccessToken

	client := &http.Client{Timeout: 10 * time.Second}

	projectList := project.List(organization, apiVersion)

	resp := vsts_api.RequestHandler(projectList, personalAccessToken, client)

	robots, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(robots, &projects)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Projects: %+v\n", projects)

	resp.Body.Close()

	// ONE PROJECT
	projectID := "884ddd52-da93-4d20-93f4-ce6b0db92812"
	project := project.Get(organization, projectID, apiVersion)

	resp1 := vsts_api.RequestHandler(project, personalAccessToken, client)

	robots1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(robots1, &singleProject)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Project: %+v\n", singleProject)

	resp1.Body.Close()

}
