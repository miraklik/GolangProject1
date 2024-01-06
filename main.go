package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const exitCommand = "!exit"

var operators = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

var ROMANS = map[string]int{
	"M": 1000, "D": 500, "C": 100,
	"L": 50, "X": 10, "V": 5, "I": 1,
}

func findOperator(line string) (string, error) {
	for op := range operators {
		if strings.Contains(line, " "+op+" ") {
			return op, nil
		}
	}
	return "", fmt.Errorf("can't find operator")
}

func validateExpression(line string) error {
	match, _ := regexp.MatchString(`^\d+[+\-*/]\d+$`, line)
	if !match {
		return fmt.Errorf("invalid expression format")
	}
	return nil
}

// Assume these functions are defined somewhere in your code
func getNumsAndType(line string, op string) (int, int, bool, error) {.
	return 0, 0, false, nil
}

func calculation(a, b int, op string) (int, error) {
	return 0, nil
}

func printResult(a, b, result int, op string, isRom bool) {
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Для выхода введите", exitCommand)
		fmt.Print("Введите пример: ")

		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == exitCommand {
			fmt.Println("Exiting..")
			return
		}

		if err := validateExpression(line); err != nil {
			fmt.Println("Error:", err)
			continue
		}

		operator, err := findOperator(line)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		a, b, isRom, err := getNumsAndType(line, operator)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		result, err := calculation(a, b, operator)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		printResult(a, b, result, operator, isRom)
	}
}
