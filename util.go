package main

import (
	"github.com/fatih/color"
)

var printError = color.New(color.Bold, color.FgRed).PrintlnFunc()
var printTitle = color.New(color.Bold, color.FgGreen).PrintfFunc()

func areSlicesEqual(s, keys []string) bool {
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
