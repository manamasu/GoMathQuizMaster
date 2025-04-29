package helper

import (
	"fmt"
	"math/rand"
	"quizmaster/internal/csvreadwriter"
	"time"
)

func TypewriterEffect(text string, delay time.Duration) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}

func GenerateMathProblemRecord() [][]int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src) //Salting (Randomization) based on our src

	numbers := make([][]int, 0, 200)

	for i := 0; i < 200; i++ {
		n1 := r.Intn(100)
		n2 := r.Intn(100)
		record := []int{n1, n2, n1 + n2}

		numbers = append(numbers, record)
	}

	return numbers
}

func GenerateCSVMathProblemsFile(filename string) {
	writer, file, err := csvreadwriter.CreateCSVWriter(filename)

	if err != nil {
		fmt.Println("Error creating CSV writer: ", err)
	}
	defer file.Close()

	headers := csvreadwriter.NewRecord([]string{"Summand1", "Summand2", "Sum"})
	csvreadwriter.WriteCSVRecord(writer, headers)

	numbers := GenerateMathProblemRecord()

	for _, v := range numbers {
		mathRecord := csvreadwriter.NewRecord(v)
		csvreadwriter.WriteCSVRecord(writer, mathRecord)
	}
}
