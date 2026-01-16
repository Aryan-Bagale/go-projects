package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// reads one full line from stdin
func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(line), nil
}

// reads and validates an integer
func readInt(reader *bufio.Reader) (int, error) {
	fmt.Print("Type a Number: ")
	input, err := readLine(reader)
	if err != nil {
		return 0, err
	}
	if input == "" {
		return 0, fmt.Errorf("empty input")
	}
	return strconv.Atoi(input)
}

// prints whether a number is even or odd
func printEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("It's Even")
	} else {
		fmt.Println("It's Odd")
	}
}

// asks user whether to continue
func askToContinue(reader *bufio.Reader) bool {
	fmt.Print("Continue? (y/n): ")
	choice, err := readLine(reader)
	if err != nil {
		return false
	}

	choice = strings.ToLower(choice)
	return choice == "y"
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		num, err := readInt(reader)
		if err != nil {
			fmt.Println("Invalid input, try again.")
			continue
		}

		printEvenOdd(num)

		if !askToContinue(reader) {
			fmt.Println("Goodbye!")
			break
		}
	}
}
