package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//Configuration for holding info
type Configuration struct {
	Origin   string
	Upstream string
}

func readConfigFile() Configuration {
	file, _ := os.Open(".spaconrc")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	return configuration
}

func writeConfigFile(c Configuration) {
	configJSON, _ := json.Marshal(c)
	err := ioutil.WriteFile(".spaconrc", configJSON, 0644)
	if err != nil {
		fmt.Errorf("%s", err)
	}
}
