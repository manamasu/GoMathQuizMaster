package main

import (
	"flag"
	"fmt"
	"quizmaster/internal/menu"
)

var filename string
var max int
var sampleSize int

func main() {

	flag.StringVar(&filename, "csv", "problems.csv", "CSV file with math problems")
	flag.IntVar(&max, "max", 100, "Max has to be positive. e.g. Numbers from 0 to max")
	flag.IntVar(&sampleSize, "size", 50, "Sample Size of math problems. Has to be at least one.")
	flag.Parse()

	if max <= 0 {
		fmt.Println("The max flag can only be used with positive numbers.")
		return
	}
	if sampleSize <= 0 {
		fmt.Println("The sample size has to be at least one to start the Quiz.")
		return
	}

	menu.MainMenu(filename, max, sampleSize)
}
