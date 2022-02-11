package functions

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// changeQuizName will display the current name of the file and will change it if required
func changeQuizName() {

	// get the current name of the file
	currentName := getCurrentName()

	// get new name
	fmt.Printf("current Name : [%s]. Please introduce the new name with csv format and press enter: ", currentName)
	newName, err := readInput()
	if err != nil {
		log.Fatalln(err)
	}
	if newName == "\n" {
		log.Fatalln("the name must not be blank")
	}

	oldPathName := fmt.Sprint("fixtures/", currentName)
	newPathName := fmt.Sprint("fixtures/", newName)

	// change name of the file
	err = os.Rename(oldPathName, newPathName)
	if err != nil {
		log.Fatalln(err)
	}

	// get the new name after user changed
	newName = getCurrentName()
	fmt.Printf("The old name was: [%s]. The new name is: [%s],", currentName, newName)
}

// getCurrentName will get the name of specified path
func getCurrentName() string {

	// get list of files in particular path
	files, err := ioutil.ReadDir("fixtures")
	if err != nil {
		log.Fatalln(err)
	}

	// it will be only one file atm
	var myFileName string
	for _, file := range files {
		myFileName = file.Name()
	}

	return myFileName
}
