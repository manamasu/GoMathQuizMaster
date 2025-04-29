package main

import (
	"fmt"
	"quizmaster/internal/csvreadwriter"
	"quizmaster/internal/helper"
	"time"
)

const filename = "problems.csv"

func main() {
	go generateCSVMathProblemsFile(filename)

	helper.TypewriterEffect("Hello to GoMathQuizmaster\n", 50*time.Millisecond)

	listMenu()
	var choice string

	_, err := fmt.Scan(&choice)

	if err != nil {
		for choice != "1" {
			fmt.Println("Please Press 1) or 2) if you want to continue")
			fmt.Scan(&choice)
		}
	}
}

func listMenu() {
	fmt.Println("Would you like to start with Addition or Subtraction?")
	fmt.Println("Press 1) Addition")
	fmt.Println("Press 2) Subtraction")
}

func generateCSVMathProblemsFile(filename string) {
	writer, file, err := csvreadwriter.CreateCSVWriter(filename)

	if err != nil {
		fmt.Println("Error creating CSV writer: ", err)
	}
	defer file.Close()

	headers := csvreadwriter.NewRecord([]string{"Summand1", "Summand2", "Sum"})
	csvreadwriter.WriteCSVRecord(writer, headers)

	numbers := helper.GenerateMathProblemRecord()

	for _, v := range numbers {
		mathRecord := csvreadwriter.NewRecord(v)
		csvreadwriter.WriteCSVRecord(writer, mathRecord)
	}
}
