package main

import "fmt"

func main() {
	const maxAttempts = 3
	password := "secret"

	for attempts := 1; attempts <= maxAttempts; attempts++ {

		fmt.Print("Enter Password: ")
		var input string
		fmt.Scan(&input)

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
