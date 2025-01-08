package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/fatih/color"
)

func HandleCSV(answers map[string]string, filename string, filepath string) {
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

func createNewData(existingCSVData [][]string, answers map[string]string) [][]string {
	var result [][]string = make([][]string, len(answers))
	var timeindex int
	keys := make([]string, len(answers))
	values := make([]string, len(answers))
	i := 0
	for k, v := range answers {
		keys[i] = k
		values[i] = v
		if values[i] == "time" {
			timeindex = i
		}
		i++
	}
	result[0] = keys
	result[1] = values

	fieldsAreEqual := reflect.DeepEqual(existingCSVData[0], keys)

	if !fieldsAreEqual {
		color.New(color.FgRed, color.Bold).Print("Your questions differ from the data that was collected until now.\nIf you proceed, all your previous data will be lost.\nContinue? [Default: 2]")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		var answer string
		for {
			fmt.Scan("%s", &answer)
			if answer == "2" || answer == "" {
				log.Fatal("Tracking data cancelled.")
			} else if answer == "1" {
				return result
			} else {
				printError("Error: Invalid input.")
				continue
			}
		}
	}

	fmt.Print(timeindex)
	//CHECKEN ob an diesem Tag schon ein Eintrag gemacht wurde, mit selben Namen
	return result
}

func saveDataToCSV(newData [][]string, filename, filepath string) {
	panic("unimplemented")
}
