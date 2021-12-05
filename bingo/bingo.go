package bingo

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/geordie/adventofcode2021/util"
)

type BingoNumbers []int
type BingoBoard [5][5]int
type BingoBoards []BingoBoard

func ParseInput(sFile string) {
	file, err := os.Open(sFile)
	if err != nil {
		log.Fatal(err)
	}

	var bingoNums BingoNumbers
	scanner := bufio.NewScanner(file)

	// Get the BingoNumbers for first row of file
	scanner.Scan()
	s := scanner.Text()
	bingoNums = bingoNums.Parse(s)
}

func (bingoNums BingoNumbers) Parse(sInput string) BingoNumbers {
	sInputs := strings.Split(sInput, ",")

	for _, elem := range sInputs {
		bingoNums = append(bingoNums, util.GetIntFromString(elem))
	}
	return bingoNums
}
