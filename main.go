package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Question struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Answers      []string `json:"answers"`
	DefaultValue string   `json:"default_value"`
	CSVName      string   `json:"csv"`
}

type ConfigData struct {
	CSVFilepath string     `json:"csv_filepath"`
	CSVFilename string     `json:"csv_filename"`
	Questions   []Question `json:"questions"`
}

func main() {
	fmt.Println("Welcome to the day tracker. Select an answer from the possible choices or just type directly if no choices are presented.\nPress Enter if you want to choose the default choice displayed in the brackets.")
	var csv_filename, csv_filepath string
	var questions []Question
	csv_filename, csv_filepath, questions = load_config()
	var selected_answers [][]string = start_questions(questions)
	// => load existing csv
	// => check if field names changed: if yes prompt user if it wants to create new one
	// => save answers to csv
	HandleCSV(selected_answers, csv_filename, csv_filepath)
}

// Load config from config.json file
// return: the csv filepath and name aswell as all the questions
func load_config() (string, string, []Question) {

	var err error
	var configBytes []byte
	configBytes, err = os.ReadFile("./config.json")
	if err != nil {
		fmt.Println("Could not load config file.")
		log.Fatal(err)
	}
	var configData ConfigData
	err = json.Unmarshal(configBytes, &configData)
	if err != nil {
		log.Fatal(err)
	}

	return configData.CSVFilename, configData.CSVFilepath, configData.Questions
}
