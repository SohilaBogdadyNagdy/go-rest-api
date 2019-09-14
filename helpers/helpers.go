package helpers

import (
	"log"
	"encoding/json"
	"io/ioutil"

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
