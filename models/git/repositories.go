package git

import (
	"github.com/hinrek/Azure-migrator/utils"
	"github.com/hinrek/Azure-migrator/vsts-api"
	"log"
	"net/http"
)

type Repository struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	URL     string `json:"url"`
	Project struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		URL   string `json:"url"`
		State string `json:"state"`
	} `json:"project"`
	RemoteURL     string `json:"remoteUrl"`
	DefaultBranch string `json:"defaultBranch,omitempty"`
}

type Repositories struct {
	Count      int          `json:"count"`
	Repository []Repository `json:"value"`
}

func (repositories *Repositories) List(organization string, project string, apiVersion string, personalAccessToken string, client *http.Client) *Repositories {
	// https://dev.azure.com/{organization}/{project}/_apis/git/repositories?api-version=5.0
	url := vsts_api.ConstructAzureUri(organization, project, "git", "repositories", apiVersion)
	httpResponse := vsts_api.ExecuteRequest("GET", url, personalAccessToken, client)

	bytes, err := utils.ReadResponseBody(*httpResponse)
	if err != nil {
		log.Fatal(err)
	}
	err = utils.JsonUnmarshal(bytes, &repositories)
	if err != nil {
		log.Fatal(err)
	}

	return repositories
}
