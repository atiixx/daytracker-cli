package main

import (
	"daytracker-cli/filemanager"
	"daytracker-cli/questions"
	"fmt"

	filemanager "github.com/atiixx/daytracker-cli"
)

func main() {
	fmt.Println("Hello, world.")
	questions.Start_questions()
	filemanager.SaveDataToCSV()
}
