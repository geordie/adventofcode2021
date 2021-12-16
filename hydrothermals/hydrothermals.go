package hydrothermals

import (
	"bufio"
	"log"
	"os"
)

func ParseInput(sFile string) {
	file, err := os.Open(sFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	// Get the BingoNumbers for first row of file
	scanner.Scan()
}
