package tests

import (
	"os"
	"reflect"
	"testing"

	"github.com/atiixx/daytracker-cli/internal/handlers"
)

func TestHandleCSV_CreatesFile(t *testing.T) {
	tempDir := t.TempDir()
	filename := "test.csv"
	filepath := tempDir + "/"
	answers := [][]string{
		{"name", "feel", "weather", "time"},
		{"Jonas", "2", "1", "2025-01-09"},
	}

	handlers.HandleCSV(answers, filename, filepath)

	// Check if file exists
	fullPath := filepath + filename
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		t.Fatalf("Expected file to be created at %s, but it does not exist", fullPath)
	}

	// Check file content
	fileContent, _ := os.ReadFile(fullPath)
	expectedContent := "name,feel,weather,time\nJonas,2,1,2025-01-09\n"
	if string(fileContent) != expectedContent {
		t.Errorf("File content does not match. Got: %s", string(fileContent))
	}
}
func TestHandleCSV_LoadsExistingFile(t *testing.T) {
	tempDir := t.TempDir()
	filename := "test.csv"
	filepath := tempDir + "/"
	initialContent := "name,feel,weather,time\nJonas,2,1,2025-01-09\n"
	os.WriteFile(filepath+filename, []byte(initialContent), 0644)

	answers := [][]string{
		{"name", "feel", "weather", "time"},
		{"Lisa", "3", "2", "2025-01-10"},
	}

	handlers.HandleCSV(answers, filename, filepath)

	fileContent, _ := os.ReadFile(filepath + filename)
	expectedContent := "name,feel,weather,time\nJonas,2,1,2025-01-09\nLisa,3,2,2025-01-10\n"
	if string(fileContent) != expectedContent {
		t.Errorf("File content does not match. Got: %s", string(fileContent))
	}
}
func TestReadExistingCSV_LoadsCorrectData(t *testing.T) {
	tempDir := t.TempDir()
	filename := "test.csv"
	filepath := tempDir + "/"
	content := "name,feel,weather,time\nJonas,2,1,2025-01-09\n"
	os.WriteFile(filepath+filename, []byte(content), 0644)

	result := handlers.ReadExistingCSV(filename, filepath)

	expected := [][]string{
		{"name", "feel", "weather", "time"},
		{"Jonas", "2", "1", "2025-01-09"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
func TestCreateNewData_AppendsAnswers(t *testing.T) {
	existingData := [][]string{
		{"name", "feel", "weather", "time"},
		{"Jonas", "2", "1", "2025-01-09"},
	}
	answers := [][]string{
		{"name", "feel", "weather", "time"},
		{"Lisa", "3", "2", "2025-01-10"},
	}

	result := handlers.CreateNewData(existingData, answers)

	expected := [][]string{
		{"name", "feel", "weather", "time"},
		{"Jonas", "2", "1", "2025-01-09"},
		{"Lisa", "3", "2", "2025-01-10"},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
func TestSaveDataToCSV_SavesCorrectly(t *testing.T) {
	tempDir := t.TempDir()
	filename := "test.csv"
	filepath := tempDir + "/"
	data := [][]string{
		{"name", "feel", "weather", "time"},
		{"Jonas", "2", "1", "2025-01-09"},
	}

	handlers.SaveDataToCSV(data, filename, filepath)

	fileContent, _ := os.ReadFile(filepath + filename)
	expectedContent := "name,feel,weather,time\nJonas,2,1,2025-01-09\n"
	if string(fileContent) != expectedContent {
		t.Errorf("File content does not match. Got: %s", string(fileContent))
	}
}
