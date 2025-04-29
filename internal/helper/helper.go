package helper

import (
	"fmt"
	"math/rand"
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
