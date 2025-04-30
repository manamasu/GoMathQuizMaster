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

func promptForOptionalInt(reader *bufio.Reader, message string) int {
	fmt.Print(message)
	userInput, _ := reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	if userInput == "" {
		return -1 // skipped
	}
	if val, err := strconv.Atoi(userInput); err == nil && val > 0 {
		return val
	}
	fmt.Println("Invalid input. Using default values.")
	return -1
}
