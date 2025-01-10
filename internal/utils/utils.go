package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/atiixx/daytracker-cli/internal/types"
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

// Load config from config.json file or create new one
// return: the csv filepath and name aswell as all the questions
func LoadConfig(config_filepath string) (string, string, []types.Question, error) {
	configBytes, err := os.ReadFile(config_filepath)
	if err != nil {
		configBytes = CreateConfig(config_filepath)
	}
	configData, err := ParseConfigData(config_filepath, configBytes)
	if err != nil {
		return "", "", nil, fmt.Errorf("invalid data inside config.json. Consider deleting config.json and start the app again to create a new default config.json. Error: %s", err)
	}
	return configData.CSVFilename, configData.CSVFilepath, configData.Questions, nil
}

func CreateConfig(config_filepath string) []byte {
	defaultConfig := map[string]interface{}{
		"csv_filepath": "./",
		"csv_filename": "daytracking.csv",
		"questions": []interface{}{
			types.Question{
				Title:        "Name",
				Answers:      []string{},
				DefaultValue: "Chuck Norris",
				CSVName:      "name",
			},
			types.Question{
				Title:        "💕 How do you feel?",
				Answers:      []string{"Very good", "Good", "Okay", "Not good", "Bad", "Very bad"},
				DefaultValue: "2",
				CSVName:      "feel",
			},
			types.Question{
				Title:        "☀️ How is the weather?",
				Answers:      []string{"Sunny", "Grey", "Rainy", "Snow", "Mixed"},
				DefaultValue: "1",
				CSVName:      "weather",
			},
			types.Question{
				Title:        "🏅 Did you exercise today?",
				Answers:      []string{"Yes", "No"},
				DefaultValue: "1",
				CSVName:      "exercise",
			},
			types.Question{
				Title:        "💤 How did you sleep last night?",
				Answers:      []string{"Good", "Okay", "Bad"},
				DefaultValue: "1",
				CSVName:      "sleep_quality",
			},
			types.Question{
				Title:        "😴 How many hours did you sleep last night?",
				Answers:      []string{"1-6 hours", "7-9 hours", "more then 9 hours"},
				DefaultValue: "2",
				CSVName:      "sleep_duration",
			},
			types.Question{
				Title:        "📗 Did you read a book?",
				Answers:      []string{"Yes", "No"},
				DefaultValue: "1",
				CSVName:      "book",
			},
			types.Question{
				Title:        "🎬 Did you watch a movie or tv show?",
				Answers:      []string{"Yes", "No"},
				DefaultValue: "1",
				CSVName:      "movie",
			},
			types.Question{
				Title:        "🎮 Did you play games?",
				Answers:      []string{"Yes", "No"},
				DefaultValue: "1",
				CSVName:      "game",
			},
			types.Question{
				Title:        "🏢 Did you do any work?",
				Answers:      []string{"Yes", "No"},
				DefaultValue: "1",
				CSVName:      "work",
			},
			types.Question{
				Title:        "💻 Did you work on a productive hobby?",
				Answers:      []string{"Yes", "No"},
				DefaultValue: "1",
				CSVName:      "productive",
			},
			types.Question{
				Title:        "🔥 Was it hot today (> 20 C)?",
				Answers:      []string{"Yes", "No"},
				DefaultValue: "1",
				CSVName:      "hot",
			},
		},
	}

	file, err := os.Create(config_filepath)
	if err != nil {
		log.Fatalf("could not create config file: %s", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(defaultConfig)
	if err != nil {
		log.Fatalf("Could not encode default config file: %s", err)
	}
	PrintTitle("\n👌 Default config.json created!\nConsider closing the app and configuring your questions first!\n")
	configBytes, err := os.ReadFile(config_filepath)
	if err != nil {
		log.Fatalf("Error: Could not read default config file under %s", config_filepath)
	}
	return configBytes
}

func ParseConfigData(config_filepath string, configBytes []byte) (types.ConfigData, error) {
	var configData types.ConfigData
	err := json.Unmarshal(configBytes, &configData)
	if err != nil {
		return types.ConfigData{}, fmt.Errorf("could not parse config file: %w", err)
	}
	switch {
	case configData.CSVFilename == "":
		return types.ConfigData{}, fmt.Errorf("missing csv filename: %w", err)
	case configData.CSVFilepath == "":
		return types.ConfigData{}, fmt.Errorf("missing csv filepath: %w", err)
	case len(configData.Questions) == 0:
		return types.ConfigData{}, fmt.Errorf("missing questions: %w", err)
	}

	return configData, nil
}
