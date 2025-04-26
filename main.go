package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const filename = "problems.csv"

type Record[T any] struct {
	record []T
}

func main() {
	fmt.Println("Hello Quizmaster")
	generateCSVMathProblemsFile(filename)
}

func generateMathProblemRecord() [][]int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src) //Salting (Randomization) based on our src

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

	headers := Record[string]{record: []string{"Summand1", "Summand2", "Sum"}}
	writeCSVRecord(writer, headers)

	numbers := generateMathProblemRecord()

	for _, v := range numbers {
		mathRecord := Record[int]{record: v}
		writeCSVRecord(writer, mathRecord)
	}
}

// Creates a file f in the current folder and passes the csv writer w through
func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {

	f, err := os.Create(filename)

	if err != nil {
		fmt.Println("There has been an error with Creating a file. Error: ", err)
		return nil, nil, err
	}
	w := csv.NewWriter(f)

	return w, f, nil
}

// writeCSVRecord writes a Record[T] to the given CSV writer.
// converts each field to a string before writing, and flushes (writes to csv).
// T can be any type that can be stringified using fmt.Sprint.
func writeCSVRecord[T any](writer *csv.Writer, rec Record[T]) {
	var stringRecord []string
	for _, v := range rec.record {
		stringRecord = append(stringRecord, fmt.Sprint(v))
	}

	err := writer.Write(stringRecord)
	if err != nil {
		fmt.Println("Error writing record to csv: ", err)
	}
	writer.Flush()
}
