package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/hinrek/Azure-migrator/models/configuration"
	"github.com/hinrek/Azure-migrator/models/project"
)

var (
	conf     configuration.Conf
	projects project.Projects
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

	req, err := http.NewRequest("GET", fmt.Sprintf("https://dev.azure.com/%s/_apis/projects?api-version=%s", organization, apiVersion), nil)

	req.Header.Add("Content-Type", `"application/json"`)
	req.SetBasicAuth("", personalAccessToken)

	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		log.Fatalf("API response: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	} else if err != nil {
		log.Fatal(err)
	}

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
}

// https://{instance}[/{team-project}]/_apis[/{area}]/{resource}?api-version={version}
// Azure DevOps Services: dev.azure.com/{organization}
// TFS: {server:port}/tfs/{collection} (the default port is 8080, and the value for collection should be DefaultCollection but can be any collection)
// resource path: The resource path is as follows: _apis/{area}/{resource}. For example _apis/wit/workitems
