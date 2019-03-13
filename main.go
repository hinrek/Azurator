package main

import (
	"fmt"
	"github.com/hinrek/Azure-migrator/models/git"
	"net/http"
	"time"

	"github.com/hinrek/Azure-migrator/models/configuration"
	"github.com/hinrek/Azure-migrator/models/project"
)

var (
	conf          configuration.Conf
	projects      project.Projects
	singleProject project.Project
	repositories  git.Repositories
)

func init() {
	configFilePath := "configs/config.yml"
	conf.Get(configFilePath)
}

func main() {
	client := &http.Client{Timeout: 10 * time.Second}

	var (
		sourceOrganization        = conf.SourceOrganization.Name
		sourceAPIVersion          = conf.SourceOrganization.APIVersion
		sourcePersonalAccessToken = conf.SourceOrganization.PersonalAccessToken

		//destinationOrganization        = conf.SourceOrganization.Name
		//destinationAPIVersion          = conf.SourceOrganization.APIVersion
		//destinationPersonalAccessToken = conf.SourceOrganization.PersonalAccessToken
	)

	// Project list
	projectList := projects.List(sourceOrganization, sourceAPIVersion, sourcePersonalAccessToken, client)

	fmt.Printf("Projects: %+v\n", projectList)

	for _, item := range projectList.Project {
		println("ITEM: ", item.Name)
	}

	// One project
	projectID := "884ddd52-da93-4d20-93f4-ce6b0db92812"
	getProject := singleProject.Get(sourceOrganization, projectID, sourceAPIVersion, sourcePersonalAccessToken, client)

	for counter, project := range projectList.Project {
		getProject := singleProject.Get(sourceOrganization, project.ID, sourceAPIVersion, sourcePersonalAccessToken, client)
		fmt.Printf("Single project: %d %+v\n", counter, getProject)
	}

	fmt.Printf("Project: %+v\n", getProject)

	// Repositories list
	repositoriesList := repositories.List(sourceOrganization, "Project%20source", sourceAPIVersion, sourcePersonalAccessToken, client)

	fmt.Printf("Repositories: %+v\n", repositoriesList)
}
