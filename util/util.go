package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/atiixx/daytracker-cli/types"
	"github.com/fatih/color"
)

var PrintError = color.New(color.Bold, color.FgRed).PrintlnFunc()
var PrintTitle = color.New(color.Bold, color.FgGreen).PrintfFunc()

func AreSlicesEqual(s, keys []string) bool {
	isEqual := true
	if len(s) != len(keys) {
		return false
	}

	for i := range s {
		if s[i] != keys[i] {
			isEqual = false
		}
	}
	return isEqual
}

// Load config from config.json file
// return: the csv filepath and name aswell as all the questions
func Load_config(config_filepath string) (string, string, []types.Question, error) {

	var err error
	var configBytes []byte
	configBytes, err = os.ReadFile(config_filepath)
	if err != nil {
		return "", "", nil, fmt.Errorf("Could not load config file: %w", err)
	}
	var configData types.ConfigData
	err = json.Unmarshal(configBytes, &configData)
	if err != nil {
		return "", "", nil, fmt.Errorf("could not parse config file: %w", err)
	}
	if configData.CSVFilename == "" {
		return "", "", nil, fmt.Errorf("missing csv filename: %w", err)
	}
	if configData.CSVFilepath == "" {

		return "", "", nil, fmt.Errorf("missing csv filepath: %w", err)
	}
	if len(configData.Questions) == 0 {

		return "", "", nil, fmt.Errorf("missing questions: %w", err)
	}

	return configData.CSVFilename, configData.CSVFilepath, configData.Questions, nil
}
