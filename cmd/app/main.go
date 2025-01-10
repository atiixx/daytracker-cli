package main

import (
	"fmt"
	"log"

	"github.com/atiixx/daytracker-cli/internal/handlers"
	"github.com/atiixx/daytracker-cli/internal/types"
	"github.com/atiixx/daytracker-cli/internal/utils"
)

func main() {
	fmt.Println("Welcome to the day tracker. Select an answer from the possible choices or just type directly if no choices are presented.\nPress Enter if you want to choose the default choice displayed in the brackets.")
	var csv_filename, csv_filepath, config_filepath string
	var questions []types.Question
	var err error
	config_filepath = "./config.json"
	csv_filename, csv_filepath, questions, err = utils.LoadConfig(config_filepath)
	if err != nil {
		log.Fatal(err)
	}
	var selected_answers [][]string = handlers.PromptUserQuestions(questions)
	handlers.HandleCSV(selected_answers, csv_filename, csv_filepath)
}
