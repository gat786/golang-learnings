package main

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type MessageStruct struct {
	Message struct {
		Content string `json:"content"`
	} `json:"message"`
}

func readTxtFile() {
	fmt.Println("\nReading Txt file :-")
	fmt.Println("\n=====================================\n")

	txtFileLocation := "files/simple.txt"
	content, err := os.ReadFile(txtFileLocation)

	if err != nil {
		fmt.Println("Error occured while reading file, stopping app")
		os.Exit(-1)
	}

	fmt.Println(string(content))
	fmt.Println("\n=====================================")
}

func readJsonFile() {
	fmt.Println("\nReading Json file :-")
	fmt.Println("\n=====================================\n")

	jsonFileLocation := "files/simple.json"
	content, err := os.ReadFile(jsonFileLocation)

	if err != nil {
		fmt.Println("Error occured while reading file, stopping app")
		os.Exit(-1)
	}

	var jsonContent MessageStruct
	jsonDecodingErr := json.Unmarshal(content, &jsonContent)
	if jsonDecodingErr != nil {
		fmt.Println("We found errors while decoding json exiting...")
		os.Exit(-1)
	}
	fmt.Println("content: ", jsonContent.Message.Content)
	fmt.Println("\n=====================================")
}

func readYamlFile() {
	fmt.Println("\nReading Yaml file :-")
	fmt.Println("\n=====================================\n")

	yamlFileLocation := "files/simple.yaml"
	content, err := os.ReadFile(yamlFileLocation)

	if err != nil {
		fmt.Println("Error occured while reading file, stopping app")
		os.Exit(-1)
	}

	var jsonContent MessageStruct

	jsonDecodingErr := yaml.Unmarshal(content, &jsonContent)
	if jsonDecodingErr != nil {
		fmt.Println("We found errors while decoding json exiting...")
		os.Exit(-1)
	}
	fmt.Println("content: ", jsonContent)
	fmt.Println("\n=====================================")
}

func main() {
	fmt.Println("Hello reading files in this one")

	readTxtFile()
	readJsonFile()
	readYamlFile()
}
