package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/atiixx/daytracker-cli/types"
	"github.com/atiixx/daytracker-cli/util"
)

func start_questions(questions []types.Question) [][]string {
	answers := make([][]string, 2)
	for i := range answers {
		answers[i] = make([]string, len(questions))
	}
	for i, q := range questions {
		fmt.Println()
		var answer string
		var choices bool = len(q.Answers) > 0

		if choices {
			util.PrintTitle("%s: \n", q.Title)
			for i, a := range q.Answers {
				fmt.Printf("%d: %s\n", i+1, a)
			}
			for {
				fmt.Printf("Choose (1-%d): [Default: %s]\n", len(q.Answers), q.DefaultValue)
				fmt.Scanf("%s", &answer)
				if answer != "" {
					number, err := strconv.Atoi(answer)
					if err != nil {
						util.PrintError("Error: Invalid input. Not a number")
						continue
					} else if number < 1 || number > len(q.Answers) {
						util.PrintError("Error: Invalid input. Out of range")
						continue
					}
					answers[0][i] = q.CSVName
					answers[1][i] = q.Answers[number-1]
				} else {
					defaultIndex, err := strconv.Atoi(q.DefaultValue)
					if err != nil {
						log.Fatalf("Invalid default value. Check config. Error: %s", err)
					}
					answers[0][i] = q.CSVName
					answers[1][i] = q.Answers[defaultIndex-1]
				}
				break
			}
		} else {
			util.PrintTitle("%s: ", q.Title)
			fmt.Printf("[Default: %s]\n", q.DefaultValue)
			fmt.Scanf("%s", &answer)
			if answer != "" {
				answers[0][i] = q.CSVName
				answers[1][i] = answer
			} else {
				answers[0][i] = q.CSVName
				answers[1][i] = q.DefaultValue
			}
		}
	}
	answers[0] = append(answers[0], "time")
	answers[1] = append(answers[1], time.Now().Format("2006-01-02"))

	return answers
}
