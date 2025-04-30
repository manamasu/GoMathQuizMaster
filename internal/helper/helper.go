package helper

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"quizmaster/internal/csvreadwriter"
	"runtime"
	"time"
)

type MathMode int

const (
	Addition MathMode = iota
	Subtraction
	Multiplication
	Division
)

func (m MathMode) String() string {
	return [...]string{"Addition", "Subtraction", "Multiplication", "Division"}[m]
}

// TypewriterEffect prints given text character by character with a delay
// simulating the look of a typewriter or animated typing effect.
func TypewriterEffect(delay time.Duration, text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(delay)
	}
}

// GenerateMathProblemRecord creates math problems based on the given MathMode and operation function.
// Returns a slice of integer slices, where each inner slice represents a single problem: [operand1, operand2, solution].
func GenerateMathProblemRecord(sampleSize int, max int, mode MathMode, mop func(int, int) int) [][]int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src) //Salting (Randomization) based on our src

	// Preallocating a slice to hold sampleSize Amount of problem records --- e.g.: default sampleSize = 100
	numbers := make([][]int, 0, sampleSize)

	switch mode {
	case Division:
		numbers = generateDivisionProblems(r, sampleSize, max)
	default:
		// For addition, subtraction, and multiplication:
		// generate random operands and calculate result by using math operation function (mop)
		for i := 0; i < sampleSize; i++ {
			a := r.Intn(max)
			b := r.Intn(max)
			record := []int{a, b, mop(a, b)}

			numbers = append(numbers, record)
		}
	}

	return numbers
}

// GenerateCSVMathProblemsFile creates a CSV file with math problems and their solutions.
// It uses the provided filename, operation mode, and math operation function (mop).
// The CSV will contain three columns with a header of: Number1, Number2, and Solution
func GenerateCSVMathProblemsFile(filename string, sampleSize int, max int, mode MathMode, mop func(int, int) int) {
	writer, file, err := csvreadwriter.CreateCSVWriter(filename)

	if err != nil {
		fmt.Println("Error creating CSV writer: ", err)
	}
	defer file.Close()

	headers := csvreadwriter.NewRecord([]string{"Number1", "Number2", "Solution"})
	csvreadwriter.WriteCSVRecord(writer, headers)

	numbers := GenerateMathProblemRecord(sampleSize, max, mode, mop)

	for _, v := range numbers {
		mathRecord := csvreadwriter.NewRecord(v)
		csvreadwriter.WriteCSVRecord(writer, mathRecord)
	}
}

// GetMathOp takes a MathMode value and returns a corresponding function
// that performs the specified mathematical operation on two integers.
func GetMathOp(mode MathMode) func(int, int) int {
	switch mode {
	case Addition:
		return func(a, b int) int { return a + b }
	case Subtraction:
		return func(a, b int) int { return a - b }
	case Multiplication:
		return func(a, b int) int { return a * b }
	case Division:
		return func(a, b int) int {
			if b == 0 {
				panic("division by zero")
			}
			return a / b
		}
	default:
		return func(a, b int) int { return 0 }
	}
}

// MapInputToMathMode maps a string input to a defined MathMode and returns the corresponding operation function
func MapInputToMathMode(input string) (MathMode, func(int, int) int, error) {
	switch input {
	case "1":
		return Addition, GetMathOp(Addition), nil
	case "2":
		return Subtraction, GetMathOp(Subtraction), nil
	case "3":
		return Multiplication, GetMathOp(Multiplication), nil
	case "4":
		return Division, GetMathOp(Division), nil
	default:
		return 0, nil, fmt.Errorf("invalid input %s", input)
	}
}

// generateDivisionProblems handles exact division math problems, avoiding division by zero
func generateDivisionProblems(r *rand.Rand, sampleSize int, max int) [][]int {
	numbers := make([][]int, 0, sampleSize)

	// For division it is important to avoid division by zero, so we need to offset r.Intn(number) by +1.
	for i := 0; i < sampleSize; i++ {
		b := r.Intn(9) + 1          // Generating a divisor between 1 and 9 (inclusive).
		result := r.Intn(max/b) + 1 // Generating a quotient between 1 and 10 (inclusive).
		a := b * result             // Calculate the dividend for exact division. e.g: b=3; result=4; a=3*4=12; 12/3=4
		record := []int{a, b, result}
		numbers = append(numbers, record)
	}

	return numbers
}

func GetOperatorSymbol(mode MathMode) string {
	switch mode {
	case Addition:
		return "+"
	case Subtraction:
		return "-"
	case Multiplication:
		return "*"
	case Division:
		return "/"
	default:
		return "?"
	}
}

// ClearTerminal allows for clearance of the terminal screen in a cross-platform way.
func ClearTerminal() {
	// Clear Terminal
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
