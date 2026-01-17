package main

import "fmt"

func main() {
	const maxAttempts = 3
	password := "secret"

	for attempts := 1; attempts <= maxAttempts; attempts++ {

		fmt.Print("Enter Password: ")
		var input string
		n, err := fmt.Scan(&input)
		if err!= nil || n != 1{
			fmt.Println("Invalid input. Try agian.")
			attempts--  // don’t count this as a used attempt
			continue
		}

		if input == password {
			fmt.Println("Password is correct")
			return
		}

		// only print "Try again" if there are remaining attempts
		if attempts < maxAttempts {
			fmt.Println("Try again")
			fmt.Println("Attempts left:", maxAttempts-attempts)
		}
	}

	fmt.Println("Locked Out")
}
