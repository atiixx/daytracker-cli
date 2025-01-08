package main

import "github.com/fatih/color"

var printError = color.New(color.Bold, color.FgRed).PrintlnFunc()
var printTitle = color.New(color.Bold, color.FgGreen).PrintfFunc()
