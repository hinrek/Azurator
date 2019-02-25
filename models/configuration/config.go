package configuration

import (
	"log"

	"github.com/hinrek/Azurator/filereader"

	yaml "gopkg.in/yaml.v2"
)

// Conf struct is for holding configuration
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

// GetConf returns configuration
func (c *Conf) GetConf() *Conf {
	// Got some help from https://stackoverflow.com/questions/30947534/reading-a-yaml-file-in-golang
	yamlFile := filereader.ReadFile("configs/config.yml")
	err := yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
