# GoMathQuizMaster

## Description

GoMathQuizMaster is a simple CLI-based math quiz application built with Go.  
It offers basic math challenges like addition, subtraction, and more.  
You can also create and load your own custom challenges using a CSV file formatted in **three columns** (Number1, Number2, Solution).

## Features

- [x] CLI-based user interface with flag options:
  - `--csv` to specify your custom problem set
  - `--max` to control the upper limit of random numbers
  - `--size` to set the number of questions
- [x] Automatically generates CSV files with random math problems
- [x] Supports basic operations:
  - [x] addition
  - [x] subtraction
  - [x] multiplication
  - [x] division
- [ ] Tests (In Progress)

  - Example:

  ```csv
  3,4,7
  10,5,15
  ```

## Getting Started

### Clone the Project

```bash
git clone https://github.com/manamasu/GoMathQuizMaster
cd GoMathQuizMaster
```

### Initialize Go module

```bash
go mod tidy
```

### Running Tests

```bash
go test ./...
```

## Run the Application

Run the CLI with default options or customize it using flags:

```bash
go run . --csv=problems.csv --max=100 --size=50
```

| Flag     | Description                                                           | Example            |
| -------- | --------------------------------------------------------------------- | ------------------ |
| `--csv`  | Path to a CSV file with math problems (3 columns: num1,num2,solution) | `--csv=custom.csv` |
| `--max`  | Upper limit for random number generation (default: 100)               | `--max=200`        |
| `--size` | Number of math problems to generate (default: 50)                     | `--size=25`        |

All flags are optional, if you skip them, the app will ask for your input or use defaults.
