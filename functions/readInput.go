package functions

import (
	"bufio"
	"os"
)

// readInput reads the input from the user
func readInput() (string, error) {
	// read stdin

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return input, nil
}
