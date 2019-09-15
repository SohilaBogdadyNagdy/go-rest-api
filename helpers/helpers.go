package helpers

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Read Json file and parsed it
func ParseJSONFile(fileName string, parsedTo interface{}) (ParsedData interface{}) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	err2 := json.Unmarshal(content, &parsedTo)
	if err2 != nil {
		log.Fatal(err2.Error())
	}
	return parsedTo
}
