package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const filename = "problems.csv"

func main() {
	fmt.Println("Hello Quizmaster")
	generateCSVMathProblemsFile(filename)
	generateMathProblemRecord()
}

func generateMathProblemRecord() [][]int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	numbers := make([][]int, 0, 200)

	for i := 0; i < 200; i++ {
		n1 := r.Intn(100)
		n2 := r.Intn(100)
		record := []int{n1, n2, n1 + n2}

		numbers = append(numbers, record)
	}

	fmt.Println(numbers)

	return numbers
}

func generateCSVMathProblemsFile(filename string) {
	writer, file, err := createCSVWriter(filename)

	if err != nil {
		fmt.Println("Error creating CSV writer: ", err)
	}
	defer file.Close()

	headers := []string{"Summand1", "Summand2", "Sum"}
	writeCSVRecord(writer, headers)
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {

	f, err := os.Create(filename)

	if err != nil {
		fmt.Println("There has been an error with Creating a file. Error: ", err)
		return nil, nil, err
	}
	w := csv.NewWriter(f)

	return w, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	if err != nil {
		fmt.Println("Error writing record to csv: ", err)
	}
}
