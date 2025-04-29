package menu

import (
	"bufio"
	"fmt"
	"quizmaster/internal/helper"
	"strconv"
	"strings"
	"time"
)

func promptYesNo(reader *bufio.Reader, prompt string) bool {
	for {
		helper.TypewriterEffect(fastEffect*time.Millisecond, prompt)
		userInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			continue
		}
		userInput = strings.TrimSpace(strings.ToLower(userInput))

		switch userInput {
		case "y":
			return true
		case "n", "exit":
			return false
		default:
			fmt.Println("Invalid input. Please enter 'y', 'n', or 'exit'.")
		}
	}
}

func promptForSampleSize(filename string, sampleSize int, max int, mode helper.MathMode, r *bufio.Reader) {
	for {
		fmt.Print("How many math problems do you want? ")
		userInput, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		userInput = strings.TrimSpace(strings.TrimSpace(userInput)) // Remove newline
		sampleSize, err := strconv.Atoi(userInput)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}

		// Generate problems
		mathOp := helper.GetMathOp(mode)
		go helper.GenerateCSVMathProblemsFile(filename, sampleSize, max, mode, mathOp)
		helper.ClearTerminal()

		break // Exit the loop once input is valid
	}
}

func promptForInt(reader *bufio.Reader, message string) int {
	for {
		fmt.Print(message)
		input, _ := reader.ReadString('\n')
		if val, err := strconv.Atoi(strings.TrimSpace(input)); err == nil {
			return val
		}
		fmt.Println("Invalid number, try again.")
	}
}
