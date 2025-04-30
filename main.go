package main

import (
	"flag"
	"fmt"
	"quizmaster/internal/helper"
	"quizmaster/internal/menu"
)

var (
	filename   string
	max        int
	sampleSize int
)

func main() {

	flag.StringVar(&filename, "csv", "problems.csv", "CSV file with math problems")
	flag.IntVar(&max, "max", helper.DefaultMax, "Max has to be positive. e.g. Numbers from 0 to max")
	flag.IntVar(&sampleSize, "size", helper.DefaultSampleSize, "Sample Size of math problems. Has to be at least one.")
	flag.Parse()

	if max <= 0 {
		fmt.Println("The max flag can only be used with positive numbers.")
		return
	}
	if sampleSize <= 0 {
		fmt.Println("The sample size has to be at least one to start the Quiz.")
		return
	}

	menu.MainMenu(filename, sampleSize, max)
}
