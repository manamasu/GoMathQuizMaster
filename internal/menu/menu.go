package menu

import (
	"bufio"
	"fmt"
	"os"
	"quizmaster/internal/helper"
	"strings"
	"time"
)

const slowEffect time.Duration = 50
const fastEffect time.Duration = 25

func MainMenu(filename string, max int, sampleSize int) {
	helper.TypewriterEffect(slowEffect*time.Millisecond, "Hello to GoMathQuizmaster\n")
	helper.TypewriterEffect(fastEffect*time.Millisecond,
		"The rules are simple — pick a math mode and solve the problem!\n"+
			"Each time you're right, your score increases.\n"+
			"Mistakes will also be tracked.\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		if !promptYesNo(reader, "Would you like to play? Enter y/n or type 'exit' to quit: ") {
			fmt.Println("See you next time, goodbye!")
			return
		}

		mode, err := selectMathMode(reader)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		fmt.Printf("You selected: %s\n", mode)

		if max == 0 {
			max = promptForInt(reader, "Enter minimum number: ")
		}

		if sampleSize == 0 {
			promptForSampleSize(filename, sampleSize, max, mode, reader)
		}
	}
}

func selectMathMode(reader *bufio.Reader) (helper.MathMode, error) {
	fmt.Println("\nChoose a math mode or type 'exit' to quit:")
	fmt.Println("1) Addition")
	fmt.Println("2) Subtraction")
	fmt.Println("3) Multiplication")
	fmt.Println("4) Division")
	fmt.Print("Enter your choice: ")

	userInput, err := reader.ReadString('\n')
	if err != nil {
		return 0, fmt.Errorf("failed to read input: %w", err)
	}

	userInput = strings.TrimSpace(strings.ToLower(userInput))

	if userInput == "exit" {
		os.Exit(0)
	}

	mode, _, err := helper.SelectMathMode(userInput)
	return mode, err
}
