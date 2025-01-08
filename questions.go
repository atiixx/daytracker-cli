package main

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

var printError = color.New(color.Bold, color.FgRed).PrintlnFunc()
var printTitle = color.New(color.Bold, color.FgGreen).PrintfFunc()

func start_questions(questions []Question) map[string]string {
	answers := make(map[string]string, len(questions))
	for _, q := range questions {
	redo:
		fmt.Println()
		var answer string
		var choices bool = len(q.Answers) > 0

		if choices {
			printTitle("%s: \n", q.Title)
			for i, a := range q.Answers {
				fmt.Printf("%d: %s\n", i+1, a)
			}
			fmt.Printf("Choose (1-%d): [Default: %s]\n", len(q.Answers), q.DefaultValue)
			fmt.Scanf("%s", &answer)
			if answer != "" {
				number, err := strconv.Atoi(answer)
				if err != nil {
					printError("Error: Invalid input. Not a number")
					goto redo
				} else if number < 1 || number > len(q.Answers) {
					printError("Error: Invalid input. Out of range")
					goto redo
				}
				answers[q.CSVName] = q.Answers[number-1]
			} else {
				answers[q.CSVName] = q.DefaultValue
			}

		} else {
			printTitle("%s: ", q.Title)
			fmt.Printf("[Default: %s]\n", q.DefaultValue)
			fmt.Scanf("%s", &answer)
			if answer != "" {
				answers[q.CSVName] = answer
			} else {
				answers[q.CSVName] = q.DefaultValue
			}

		}

	}
	return answers
}
