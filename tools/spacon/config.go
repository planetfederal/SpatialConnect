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

//ProjectCommands for project.json
type ProjectCommands struct {
	Build           string
	Test            string `json:"unit_test"`
	Run             string
	Deploy          string
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
		fmt.Printf("%s", err)
	}
}

func readProjectConfigFile(r Repo) ProjectCommands {
	file, _ := os.Open(r.name + "/project.json")
	decoder := json.NewDecoder(file)
	commands := ProjectCommands{}
	err := decoder.Decode(&commands)
	if err != nil {
		fmt.Println("error:", err)
	}
	return commands
}
