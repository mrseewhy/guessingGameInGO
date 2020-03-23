package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Guess the number?")

	source := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10)

	var guess int

	for {
		fmt.Println("Please input your guess: ")
		fmt.Scan(&guess)

		if guess > secretNumber {
			fmt.Println("Too Big")
		} else if guess < secretNumber {
			fmt.Println("Too Small")
		} else {
			fmt.Println("Correct! You Win :)")
			break
		}
	}
}
