package main

import (
	"fmt"
	"math/rand"
)

var d, n, numb, att int

func main() {
	// Initialization of code

	fmt.Println("Welcome to my random number generator guessing game!\nPlease select your difficulty: ")
	
	// Choosing difficulty

BEGINNING:
	fmt.Println("Difficulty:\n Easy (Press 1)\n Medium (Press 2) \n Hard (Press 3)")
	fmt.Scan(&d) 
	switch {
	case d == 1:
		fmt.Println("You chose the Easy difficulty!")
		numb = rand.Intn(10)
	case d == 2:
		fmt.Println("You chose the Medium difficulty!")
		numb = rand.Intn(25)
	case d == 3:
		fmt.Println("You chose the Hard difficulty!")
		numb = rand.Intn(50)
	default:
		fmt.Println("Invalid choice! Please try again!")
		goto BEGINNING
	}

	// Main game code
	
	fmt.Println("A random number was generated!")

	fmt.Println("Please try to guess the number!")

MAIN_STAGE:
	fmt.Scan(&n)
	if n != numb {
		fmt.Println("Your guess was incorrect! Please try again!")
		if numb < n {
			fmt.Println("The number is smaller.")
		} else {
			fmt.Println("The number is larger.")
		}
		att++
		goto MAIN_STAGE
	} else {
		fmt.Println("You guessed the number correctly!")
		att++
		goto END
	}

	// End celebration

END:	
	fmt.Println("You won the game! The number of attempts is:", att)	
}