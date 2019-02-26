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

	DestinationOrganization struct {
		Name                string `yaml:"name"`
		PersonalAccessToken string `yaml:"personalAccessToken"`
		APIVersion          string `yaml:"apiVersion"`
	} `yaml:"destinationOrganization"`
}

// GetConf returns configuration
func (c *Conf) GetConf() *Conf {
	return c.setConf()
}

func (c *Conf) setConf() *Conf {
	bytes := filereader.ReadFile("configs/config.yml")
	err := yaml.Unmarshal(bytes, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}