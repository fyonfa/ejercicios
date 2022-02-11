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

// executeInput selects the action to do from the user
func executeInput() {
	frontPlate()
	userInput, _ := readInput()
	fmt.Println("====", userInput)

	switch userInput {
	case changeNameCsv:
		changeQuizName()
	case numberOfQuestions:
		displayNumberOfQuestions()
	case play:
		playGame()
	case quit:
		log.Fatalln("Has salido del juego")
	default:
		fmt.Println("not value inputted")
	}
}
