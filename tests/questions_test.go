package tests

import (
	"testing"

	"github.com/atiixx/daytracker-cli/internal/handlers"
	"github.com/atiixx/daytracker-cli/internal/types"
)

var mockQuestions = []types.Question{
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

func TestQuestions(t *testing.T) {
	question := mockQuestions[0]
	number, err := handlers.ParseChoiceAnswer("1", question)
	if err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}
	answerstring := question.Answers[number]
	if answerstring != "Happy" {
		t.Fatalf("Expected Happy or Sunny as the Answer but got: %s", answerstring)
	}
	question = mockQuestions[1]
	number, err = handlers.ParseChoiceAnswer("1", question)
	if err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}
	answerstring = question.Answers[number]
	if answerstring != "Sunny" {
		t.Fatalf("Expected Sunny as the Answer but got: %s", answerstring)
	}
	number, err = handlers.ParseChoiceAnswer("3", question)
	if err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}
	answerstring = question.Answers[number]
	if answerstring != "Cloudy" {
		t.Fatalf("Expected Cloudy as the Answer but got: %s", answerstring)
	}

}

func TestQuestions_DefaultValue(t *testing.T) {
	question := mockQuestions[0]
	number, err := handlers.ParseChoiceAnswer("", question)
	if err != nil {
		t.Fatalf("expected no error but got: %s", err)
	}
	answerstring := question.Answers[number]
	if answerstring != "Sad" {
		t.Fatalf("Expected Defaultvalue Sad as the Answer but got: %s", answerstring)
	}
}

func TestQuestions_InvalidInput(t *testing.T) {
	question := mockQuestions[0]
	_, err := handlers.ParseChoiceAnswer("0", question)
	if err == nil {
		t.Fatalf("expected an out of range error but got nil")
	}
	_, err = handlers.ParseChoiceAnswer("4", question)
	if err == nil {
		t.Fatalf("expected an out of range error but got nil")
	}
	_, err = handlers.ParseChoiceAnswer("asd", question)
	if err == nil {
		t.Fatalf("expected an not a number error but got nil")
	}
}
