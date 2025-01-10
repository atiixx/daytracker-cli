package tests

import (
	"testing"

	"github.com/atiixx/daytracker-cli/internal/handlers"
	"github.com/atiixx/daytracker-cli/internal/types"
)

func TestQuestions(t *testing.T) {
	mockQuestions := []types.Question{
		{
			Title:        "What is your mood?",
			CSVName:      "feel",
			Answers:      []string{"Happy", "Sad", "Neutral"},
			DefaultValue: "2",
		},
		{
			Title:        "What's the weather like?",
			CSVName:      "weather",
			Answers:      []string{"Sunny", "Rainy", "Cloudy"},
			DefaultValue: "3",
		},
	}

	answer := "1"
	for i := range mockQuestions {
		number, err := handlers.ParseChoiceAnswer(answer, mockQuestions[i])
		if err != nil {
			t.Fatalf("expected no error but got: %s", err)
		}
		answerstring := mockQuestions[i].Answers[number]
		if answerstring != "Happy" && answerstring != "Sunny" {
			t.Fatalf("Expected Happy or Sunny as the Answer but got: %s", answerstring)
		}
	}
}
