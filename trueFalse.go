package main

import "fmt"

func main() {

	fmt.Println("Enter an integer value: ")

	compare()
}

func compare() {
	secretValue := 88
	var guess int
	fmt.Scanln(&guess)
	if guess == secretValue {
		fmt.Println(" You got it right!")
	} else if guess > secretValue {
		fmt.Println("Too high, try again next time!")
	} else if guess < secretValue {
		fmt.Println("Too low, try again next time!")
	}
}
