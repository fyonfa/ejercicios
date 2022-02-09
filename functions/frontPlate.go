package functions

import (
	"fmt"
)

// FrontPlate will generate a front plate of options to the user
func FrontPlate() {
	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++\n" +
		"Hi, welcome to the quiz. this are the options, press enter after selection:\n" +
		"1 - to change name of folder\n" +
		"2 - to see how many questions the quiz have\n" +
		"3 - to play\n" +
		"4 - to end the program\n" +
		"++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}
