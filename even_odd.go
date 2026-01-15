package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var choice string

	for {
		fmt.Print("Type a Number: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Try to convert the whole input to an integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, try again.")
			continue
		}

		if num%2 == 0 {
			fmt.Println("It's Even")
		} else {
			fmt.Println("It's Odd")
		}

		fmt.Print("Continue? (y/n): ")
		choice, _ = reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		if choice != "y" {
			fmt.Println("Goodbye!")
			break
		}
	}
}
