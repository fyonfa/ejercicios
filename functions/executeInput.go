package functions

import (
	"fmt"
	"log"
)

const (
	changeNameCsv     = "1\n"
	numberOfQuestions = "2\n"
	play              = "3\n"
	quit              = "4\n"
)

// ExecuteInput selects the action to do from the user
func ExecuteInput() {
	userInput, _ := readInput()
	fmt.Println("====", userInput)

	switch userInput {
	case changeNameCsv:
		fmt.Println("this is change name csv")

	case numberOfQuestions:
		fmt.Println("this is number of questions")

	case play:
		fmt.Println("you have chosen to play")

	case quit:
		log.Fatalln("Has salido del juego")
	default:
		fmt.Println("not value inputted")
	}
}
