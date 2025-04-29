# GoMathQuizMaster

## Description

GoMathQuizMaster is a simple CLI-based math quiz application built with Go.  
It offers basic math challenges like addition, subtraction, and more.  
You can also create and load your own custom challenges using a CSV file formatted in **three columns** (Number1, Number2, Solution).

## Features

- [] CLI-based user interface with flag options
- [x] Automatically generates CSV files with random math problems
- [x] Supports basic operations:
  - [x] addition
  - [x] subtraction
  - [x] multiplication
  - [x] division
- [] Support for custom challenges via CSV files (3 columns needed: Number1, Number2, Solution).
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
