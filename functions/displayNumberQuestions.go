package functions

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// displayNumberOfQuestions display the user how many questions the file has
// counting as every line like new question
func displayNumberOfQuestions() {
	fileName := fmt.Sprint("fixtures/", getCurrentName())
	fmt.Println(fileName)
	// open the file
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	//for {
	//	lines, err := r.Read()
	//	if err == io.EOF {
	//		break
	//	}
	//	fmt.Println(lines)
	//	r.ReadAll()
	//}
	records, _ := r.ReadAll()
	// fmt.Println("This are the questions:",records)
	fmt.Println("The number of questions are: ", len(records))
}
