package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hinrek/Azure-migrator/models/configuration"
	"github.com/hinrek/Azure-migrator/models/project"
)

var (
	conf               configuration.Conf
	projects      project.Projects
	singleProject project.Project
	//repositories  git.Repositories
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

		destinationOrganization        = conf.DestinationOrganization.Name
		destinationAPIVersion          = conf.DestinationOrganization.APIVersion
		destinationPersonalAccessToken = conf.DestinationOrganization.PersonalAccessToken
	)

	log.Println("GET Project List")
	projectList := projects.List(sourceOrganization, sourceAPIVersion, sourcePersonalAccessToken, client)

	//log.Println("GET Single Project")
	//projectID := "726e343e-b03d-4633-85de-6642d4b850fa"
	//getProject := singleProject.Get(sourceOrganization, projectID, sourceAPIVersion, sourcePersonalAccessToken, client)

	//log.Printf("Project: %+v\n", getProject)

	//log.Println("POST Single project")
	//postProject := getProject.Create(destinationOrganization, destinationAPIVersion, destinationPersonalAccessToken, client)
	//log.Println(postProject.Body)

	log.Println("POST All Projects")
	for _, item := range projectList.Project {
		log.Println("Project: ", item.Name)

		getProject 
		
  := singleProject.Get(sourceOrganization, item.ID, sourceAPIVersion, sourcePersonalAccessToken, client)
		getProject.Create(destinationOrganization, destinationAPIVersion, destinationPersonalAccessToken, client)
	}

	//log.Println("GET Repositories List")
	//repositoriesList := repositories.List(sourceOrganization, "Project%20source", sourceAPIVersion, sourcePersonalAccessToken, client)

	//log.Printf("Repositories: %+v\n", repositoriesList)
}
