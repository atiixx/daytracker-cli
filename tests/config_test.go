package tests

import (
	"os"
	"testing"

	"github.com/atiixx/daytracker-cli/internal/utils"
)

//Test if config gets loaded and has the right amount of questions, filename and filepath.
//Should catch empty filepath, filename or questions
//Test if it gets handled when no config gets loaded

//Test if default values are being applied
//Test if answers get applied
//Test if invalid inputs get handled

// Test if CSV File is getting created at filepath with filename when none exists
// Test if CSV File gets loaded when one exists
// Test if CSV File data gets loaded correctly
// Test if answers get append correctly to csv data
// Test if CSV File gets saved correctly at the same place as filepath and filename
// Test if CSV File overrides existing one
var configContent string = `{
  "csv_filepath": "/tmp",
  "csv_filename": "testfile.csv",
  "questions": [
    {
      "title": "Name",
      "answers": [],
      "default_value": "Jonas",
      "csv": "name"
    },
    {
      "title": "How do you feel?",
      "answers": ["Very good", "Good", "Okay", "Not good", "Bad", "Very bad"],
      "default_value": "2",
      "csv": "feel"
    },
    {
      "title": "How is the weather?",
      "answers": ["Sunny", "Grey", "Rainy", "Snow", "Mixed"],
      "default_value": "1",
      "csv": "weather"
    }]}`
var missingFilepathConfig string = `{
  "csv_filepath": "",
  "csv_filename": "testfile.csv",
  "questions": [
    {
      "title": "Name",
      "answers": [],
      "default_value": "Jonas",
      "csv": "name"
    }]}`

var missingFilenameConfig string = `{
  "csv_filepath": "/tmp",
  "csv_filename": "",
  "questions": [
    {
      "title": "Name",
      "answers": [],
      "default_value": "Jonas",
      "csv": "name"
    }]}`

var missingQuestionsConfig string = `{
  "csv_filepath": "/tmp",
  "csv_filename": "testfile.csv",
  "questions": []}`

var emptyConfig string = `{}`

func TestLoadConfig(t *testing.T) {
	tempFile := "test_config.json"
	err := os.WriteFile(tempFile, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	defer os.Remove(tempFile)
	filename, filepath, questions, err := utils.LoadConfig(tempFile)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if filename != "testfile.csv" {
		t.Errorf("Expected csv_filename 'testfile.csv', got '%s'", filename)
	}
	if filepath != "/tmp" {
		t.Errorf("Expected csv_filepath '/tmp', got '%s'", filepath)
	}
	if len(questions) != 3 {
		t.Errorf("Expected 3 questions, got %d", len(questions))
	}
	if questions[1].Title != "How do you feel?" {
		t.Errorf("Expected question title 'How do you feel?', got '%s'", questions[0].Title)
	}
}
func TestLoadConfig_InvalidFile(t *testing.T) {
	// Call the function with a non-existent file
	_, _, _, err := utils.LoadConfig("nonexistent.json")
	if err == nil {
		t.Fatal("Expected an error for missing file, but got nil")
	}
}

func TestLoadConfig_InvalidJSON(t *testing.T) {
	tempFile := "test_invalid_config.json"
	err := os.WriteFile(tempFile, []byte("{invalid_json}"), 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	defer os.Remove(tempFile)
	_, _, _, err = utils.LoadConfig(tempFile)
	if err == nil {
		t.Fatal("Expected an error for invalid JSON, but got nil")
	}
}
func TestLoadingConfig_MissingFilepath(t *testing.T) {
	tempFile := "test_config.json"
	err := os.WriteFile(tempFile, []byte(missingFilepathConfig), 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	defer os.Remove(tempFile)
	_, _, _, err = utils.LoadConfig(tempFile)
	if err == nil {
		t.Fatalf("Expected an Error for missing filepath but got nil")
	}
}

func TestLoadingConfig_MissingFilename(t *testing.T) {
	tempFile := "test_config.json"
	err := os.WriteFile(tempFile, []byte(missingFilenameConfig), 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	defer os.Remove(tempFile)
	_, _, _, err = utils.LoadConfig(tempFile)
	if err == nil {
		t.Fatalf("Expected an Error for missing filename but got nil")
	}
}

func TestLoadingConfig_MissingQuestions(t *testing.T) {
	tempFile := "test_config.json"
	err := os.WriteFile(tempFile, []byte(missingQuestionsConfig), 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	defer os.Remove(tempFile)
	_, _, _, err = utils.LoadConfig(tempFile)
	if err == nil {
		t.Fatalf("Expected an Error for missing questions but got nil")
	}
}

func TestLoadingConfig_EmptyConfig(t *testing.T) {
	tempFile := "test_config.json"
	err := os.WriteFile(tempFile, []byte(emptyConfig), 0644)
	if err != nil {
		t.Fatalf("Could not create test file: %v", err)
	}
	defer os.Remove(tempFile)
	_, _, _, err = utils.LoadConfig(tempFile)
	if err == nil {
		t.Fatalf("Expected an Error for missing configdata but got nil")
	}
}
