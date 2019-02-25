package filereader

import (
	"io/ioutil"
	"log"
)

func ReadFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Reading file error: #%v", err)
	}
	return content
}
