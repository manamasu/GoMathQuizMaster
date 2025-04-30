package menu

import (
	"bufio"
	"fmt"
	"os"
	"quizmaster/internal/csvreadwriter"
	"quizmaster/internal/helper"
	"strconv"
	"strings"
	"time"
)

const slowEffect time.Duration = 50
const fastEffect time.Duration = 25

var (
	score    int
	mistakes int
)

func MainMenu(filename string, sampleSize int, max int) {
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

		helper.ClearTerminal()

		fmt.Printf("You selected: %s\n", mode)

		if max == helper.DefaultMax {
			fmt.Printf("default Max is: %v, default Samplesize is: %v", helper.DefaultMax, helper.DefaultSampleSize)
			if input := promptForOptionalInt(reader, "\nHow many math problems do you want?\n"+
				"Enter a max number from 1 to max (or press Enter to keep default 100): "); input != -1 {
				max = input
			}
		}

		if sampleSize == helper.DefaultSampleSize {
			if input := promptForOptionalInt(reader, "Enter number of problems (or press Enter to keep default 50): "); input != -1 {
				sampleSize = input
			}
		}

		helper.GenerateCSVMathProblemsFile(filename, sampleSize, max, mode, helper.GetMathOp(mode))
		helper.ClearTerminal()

		records, err := csvreadwriter.ReadCSVRecords(filename)
		if err != nil {
			fmt.Println("Error loading problems:", err)
			return
		}
		startQuiz(records, score, mistakes, mode, reader)
		displayScore()
		break
	}
}

func startQuiz(records []csvreadwriter.Record[int], score int, mistakes int, mode helper.MathMode, reader *bufio.Reader) {
	operator := helper.GetOperatorSymbol(mode)
	for i, r := range records {
		nums := r.Record
		if len(nums) < 3 {
			continue
		}
		num1, num2, solution := nums[0], nums[1], nums[2]

		fmt.Printf("Math-Problem %d: %d %s %d = ", i+1, num1, operator, num2)

		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		userAnswer, err := strconv.Atoi(userInput)

		if err != nil || userAnswer != solution {
			fmt.Printf("%s is Incorrect.\n", userInput)
			mistakes++
		} else {
			fmt.Printf("%s is Correct!\n", userInput)
			score++
		}
	}
}

func displayScore() {
	fmt.Printf("Congratulations! You had a score of: %v and did %v mistakes.\n", score, mistakes)
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
