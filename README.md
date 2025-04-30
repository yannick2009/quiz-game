# Quiz Game

## Description

Quiz Game is a simple command-line application that allows users to answer math quizzes under time constraint. The application reads questions and answers from a CSV file and verifies user responses in real-time.

## Features

- Reads questions from a built-in or custom CSV file
- Configurable timer to limit response time
- ASCII art interface
- Real-time answer validation
- Calculation and display of final score

## Installation

```bash
git clone https://github.com/your-username/quiz-game.git
cd quiz-game
go mod tidy
go build
```

## Usage

Run the application without parameters to use the default CSV file and a 30-second time limit:

```bash
./quiz-game
```

### Available Options

- `-path`: Specify a path to a custom CSV file
- `-limit`: Set the time limit in seconds to complete the quiz

Example with parameters:

```bash
./quiz-game -path=my-quiz.csv -limit=60
```

## CSV File Format

The CSV file must be formatted with two columns:
- First column contains the question
- Second column contains the correct answer

Example:
```
5+5,10
7+3,10
1+1,2
```

## Learning Objectives

- File handling in Go
- Using flags for CLI parameters
- Managing timers and goroutines
- Integrating static files with the `embed` directive
- Processing standard input/output

## License

This project is licensed under the MIT License. See the LICENSE file for details.
