package filereader

import (
	"io/ioutil"
	"log"
)

func ReadFile(filePath string) []byte {
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Reading file error: #%v", err)
	}
	return yamlFile
}
