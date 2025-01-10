package handlers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/atiixx/daytracker-cli/internal/utils"
	"github.com/fatih/color"
)

func HandleCSV(answers [][]string, filename string, filepath string) {
	var existingCSVData [][]string = readExistingCSV(filename, filepath)
	var newData [][]string = createNewData(existingCSVData, answers)
	saveDataToCSV(newData, filename, filepath)
}

func readExistingCSV(filename string, filepath string) [][]string {
	csvBytes, err := os.ReadFile(filepath + filename)
	if err != nil {
		var emptyData [][]string
		return emptyData
	}
	r := csv.NewReader(strings.NewReader(string(csvBytes)))
	csvData, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return csvData
}

func createNewData(existingCSVData [][]string, answers [][]string) [][]string {

	if len(existingCSVData) == 0 {
		return answers
	}
	fieldsAreEqual := utils.AreSlicesEqual(existingCSVData[0], answers[0])

	if !fieldsAreEqual {
		color.New(color.FgRed, color.Bold).Println("Your questions differ from the data that was collected until now.\nIf you proceed, all your previous data will be lost.\nContinue? [Default: 2]")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		var answer string
		for {
			fmt.Scanf("%s", &answer)
			if answer == "2" || answer == "" {
				log.Fatal("Tracking data cancelled.")
			} else if answer == "1" {
				return answers
			} else {
				utils.PrintError("Error: Invalid input.")
				continue
			}
		}
	}

	if answers[1][len(answers[0])-1] == existingCSVData[len(existingCSVData)-1][len(existingCSVData[0])-1] {
		color.New(color.FgRed, color.Bold).Println("You already have an entry for today.\nDo you want to override it? [Default: 2]")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		var answer string
		for {
			fmt.Scanf("%s", &answer)
			if answer == "2" || answer == "" {
				log.Fatal("Tracking data cancelled.")
			} else if answer == "1" {
				existingCSVData[len(existingCSVData)-1] = answers[1]
				return existingCSVData
			} else {
				utils.PrintError("Error: Invalid input.")
				continue
			}
		}
	}
	return append(existingCSVData, answers[1])
}

func saveDataToCSV(newData [][]string, filename, filepath string) {
	f, err := os.Create(filepath + filename)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(f)
	writer.WriteAll(newData)
	writer.Flush()
	if err := writer.Error(); err != nil {
		log.Fatal(err)
	}
}
