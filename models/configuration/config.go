package configuration

import (
	"github.com/hinrek/Azure-migrator/filereader"
	"log"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	SourceOrganization struct {
		Name                string `yaml:"name"`
		PersonalAccessToken string `yaml:"personalAccessToken"`
		APIVersion          string `yaml:"apiVersion"`
	} `yaml:"sourceOrganization"`

	DestinationOrganization struct {
		Name                string `yaml:"name"`
		PersonalAccessToken string `yaml:"personalAccessToken"`
		APIVersion          string `yaml:"apiVersion"`
	} `yaml:"destinationOrganization"`
}

func (c *Conf) Configuration(path string) *Conf {
	bytes, err := filereader.ReadFile(path)
	if err != nil {
		log.Fatalf("Filereader: %v", err)
	}

	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
