package main

import (
	"fmt"
	"log"

	"github.com/atiixx/daytracker-cli/types"
	"github.com/atiixx/daytracker-cli/util"
)

func main() {
	fmt.Println("Welcome to the day tracker. Select an answer from the possible choices or just type directly if no choices are presented.\nPress Enter if you want to choose the default choice displayed in the brackets.")
	var csv_filename, csv_filepath, config_filepath string
	var questions []types.Question
	var err error
	config_filepath = "./config.json"
	csv_filename, csv_filepath, questions, err = util.Load_config(config_filepath)
	if err != nil {
		log.Fatal(err)
	}
	var selected_answers [][]string = start_questions(questions)
	HandleCSV(selected_answers, csv_filename, csv_filepath)
}
