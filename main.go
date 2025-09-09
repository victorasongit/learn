package main

import (
	"fmt"
	"math/rand"
)

var d, n, numb, att, c int
var s string

func main() {
	// Initialization of code
MENU:
	fmt.Println("\nWelcome to my random number generator guessing game!")
	fmt.Println("\nThis is the main menu. Please choose your desired option: ")
	fmt.Println("1. Start the game \n2. Read the rules of the game \n3. Quit the game")

	fmt.Scan(&c)
	switch {
	case c == 1:
		goto BEGINNING
	case c == 2:
		fmt.Println("The rules of the game are simple: the computer generates a random number and the user ( i.e. you ) will have to guess it.\nGood luck!")
		fmt.Println("Would you like to try the game now? Please type Y or N")
		fmt.Scan(&s)
		switch {
		case s == "Y":
			goto BEGINNING
		case s == "N":
			goto MENU
		}
	case c == 3:
		goto QUIT
	}

	// Choosing difficulty
	s = ""
BEGINNING:
	fmt.Println("\nDifficulty:\nEasy (Press 1)\nMedium (Press 2) \nHard (Press 3)")
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
GUESS_AGAIN:
	fmt.Scan(&n)
	if n != numb {
		fmt.Println("Your guess was incorrect! Please try again!")
		if numb < n {
			fmt.Println("The number is smaller.")
		} else {
			fmt.Println("The number is larger.")
		}
		att++
		goto GUESS_AGAIN
	} else {
		fmt.Println("You guessed the number correctly!")
		att++
		goto END
	}

	// End celebration

END:
	fmt.Println("You won the game! The number of attempts is:", att)
	fmt.Println("Would you like to play again? Please type Y or N.")
	fmt.Scan(&s)
	switch {
	case s == "Y":
		att = 0
		goto BEGINNING
	case s == "N":
		goto QUIT
	}
QUIT:
	fmt.Println("Goodbye!")
}
