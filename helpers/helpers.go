package helpers

import (
	"api/models"
	"encoding/json"
	"io/ioutil"
	"log"
)

// Read Json file and parsed it
func ParseJSONFile(filePath string, parsedTo interface{}) (ParsedData interface{}) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	err2 := json.Unmarshal(content, &parsedTo)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	return parsedTo
}

func GetConfigurations() models.Config {
	content, err := ioutil.ReadFile("configs/config.json")
	if err != nil {
		log.Fatal(err)
	}

	var config models.Config
	err2 := json.Unmarshal(content, &config)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	return config
}
