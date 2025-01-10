package handlers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/atiixx/daytracker-cli/internal/types"
	"github.com/atiixx/daytracker-cli/internal/utils"
)

func PromptUserQuestions(questions []types.Question) [][]string {
	answers := make([][]string, 2)
	for i := range answers {
		answers[i] = make([]string, len(questions))
	}
	for i, q := range questions {
		fmt.Println()
		var answer string
		var choices bool = len(q.Answers) > 0

		if choices {
			utils.PrintTitle("%s: \n", q.Title)
			for i, a := range q.Answers {
				fmt.Printf("%d: %s\n", i+1, a)
			}
			for {
				fmt.Printf("Choose (1-%d): [Default: %s]\n", len(q.Answers), q.DefaultValue)
				fmt.Scanf("%s", &answer)
				number, err := ParseChoiceAnswer(answer, q)
				if err != nil {
					utils.PrintError(err)
					continue
				}
				answers[0][i] = q.CSVName
				answers[1][i] = q.Answers[number]
				break
			}
		} else {
			utils.PrintTitle("%s: ", q.Title)
			fmt.Printf("[Default: %s]\n", q.DefaultValue)
			fmt.Scanf("%s", &answer)
			answers[0][i] = q.CSVName
			if answer != "" {
				answers[1][i] = answer
			} else {
				answers[1][i] = q.DefaultValue
			}
		}
	}
	answers[0] = append(answers[0], "time")
	answers[1] = append(answers[1], time.Now().Format("2006-01-02"))

	return answers
}

func ParseChoiceAnswer(answer string, q types.Question) (int, error) {
	if answer != "" {
		number, err := strconv.Atoi(answer)
		if err != nil {
			return 0, fmt.Errorf("invalid Input: not a number. %s", err)
		} else if number < 1 || number > len(q.Answers) {

			return 0, fmt.Errorf("invalid Input: out of range")

		}
		return number - 1, nil
	} else {
		defaultIndex, err := strconv.Atoi(q.DefaultValue)
		if err != nil {
			return 0, fmt.Errorf("invalid default value. Check config. Error: %s", err)
		}
		return defaultIndex - 1, nil
	}
}
