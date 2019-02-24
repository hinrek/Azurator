package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	SourceOrganization struct {
		Name                string `yaml:"name"`
		PersonalAccessToken string `yaml:"personalAccessToken"`
		APIVersion          string `yaml:"apiVersion"`
	} `yaml:"sourceOrganization"`

	Destinationrganization struct {
		Name                string `yaml:"name"`
		PersonalAccessToken string `yaml:"personalAccessToken"`
		APIVersion          string `yaml:"apiVersion"`
	} `yaml:"destinationrganization"`
}

func main() {

	type Project struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		URL            string    `json:"url"`
		State          string    `json:"state"`
		Revision       int       `json:"revision"`
		Visibility     string    `json:"visibility"`
		LastUpdateTime time.Time `json:"lastUpdateTime"`
		Description    string    `json:"description,omitempty"`
	}

	type Projects struct {
		Count   int       `json:"count"`
		Project []Project `json:"value"`
	}

	var projects Projects
	var conf Conf

	yamlFile, err := ioutil.ReadFile("configs/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(conf)

	organization := conf.SourceOrganization.Name
	apiVersion := conf.SourceOrganization.APIVersion
	personalAccessToken := conf.SourceOrganization.PersonalAccessToken

	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://dev.azure.com/%s/_apis/projects?api-version=%s", organization, apiVersion), nil)

	req.Header.Add("Content-Type", `"application/json"`)
	req.SetBasicAuth("", personalAccessToken)

	resp, err := client.Do(req)
	if err != nil {
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
