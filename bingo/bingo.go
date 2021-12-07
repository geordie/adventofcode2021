package bingo

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/geordie/adventofcode2021/util"
)

type BingoNumbers []int
type BingoBoard [][]int
type BingoBoards []BingoBoard

func ParseInput(sFile string) (BingoNumbers, BingoBoards) {
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

	var bingoBoards BingoBoards
	var bingoBoard BingoBoard

	i := 0
	// Get BingoBoards in chunks of 6 lines, 1 blank, then 5 with numbers
	for scanner.Scan() {
		// On blank rows, create a new BingoBoard
		if len(scanner.Text()) == 0 {
			i = 0
			bingoBoard = BingoBoard{}
			continue
		} else {
			sRow := scanner.Text()
			arrRow := parseRow(sRow)
			bingoBoard = append(bingoBoard, arrRow)
			i++
			if i == 5 {
				bingoBoards = append(bingoBoards, bingoBoard)
			}
		}
	}

	return bingoNums, bingoBoards
}

// Runs numbers on a set of bingo boards and returns the first winning board
func (bingoBoards BingoBoards) RunGame(bingoNums BingoNumbers) (BingoBoard, int) {
	var winningBoard BingoBoard
	lastNumCalled := 0

	for _, elem := range bingoNums {
		bingoBoards.applyNumberToBoards(elem)
		possibleWinner := bingoBoards.findWinner()

		if possibleWinner.Value() > 0 {
			winningBoard = possibleWinner
			lastNumCalled = elem
			break
		}
	}

	return winningBoard, winningBoard.Value() * lastNumCalled
}

func (bingoBoards BingoBoards) applyNumberToBoards(number int) {
	for _, bingoBoard := range bingoBoards {
		bingoBoard.applyNumberToBoard(number)
	}
}

func (bingoBoards BingoBoards) findWinner() BingoBoard {
	var winner BingoBoard

	for _, bingoBoard := range bingoBoards {
		if bingoBoard.isWinner() {
			winner = bingoBoard
			break
		}
	}

	return winner
}

func (bingoBoard BingoBoard) applyNumberToBoard(number int) {
	for i, row := range bingoBoard {
		for j, cell := range row {
			if number == cell {
				bingoBoard[i][j] = -1
			}
		}
	}
}

func (bingoBoard BingoBoard) isWinner() bool {
	isWinner := false

	columnTracker := [5]int{0, 0, 0, 0, 0}

	for _, row := range bingoBoard {
		iInARow := 0
		for j, cell := range row {
			if cell == -1 {
				iInARow++
				columnTracker[j]++
				if iInARow == 5 {
					return true
				}
			}
		}
	}

	for _, col := range columnTracker {
		if col == 5 {
			return true
		}
	}

	return isWinner
}

func (bingoBoard BingoBoard) Value() int {
	iResult := 0

	for _, row := range bingoBoard {
		for _, cell := range row {
			if cell > 0 {
				iResult += cell
			}
		}
	}

	return iResult
}

func (bingoBoard BingoBoard) String() string {
	var sb strings.Builder
	sb.WriteString("--------------------\n")
	for _, row := range bingoBoard {
		for _, cell := range row {
			sb.WriteString(" " + strconv.Itoa(cell))
		}
		sb.WriteString("\n")
	}
	sb.WriteString("--------------------\n")
	return sb.String()
}

func (bingoNums BingoNumbers) Parse(sInput string) BingoNumbers {
	sInputs := strings.Split(sInput, ",")

	for _, elem := range sInputs {
		bingoNums = append(bingoNums, util.GetIntFromString(elem))
	}
	return bingoNums
}

func parseRow(sRow string) []int {
	arrSRow := strings.Split(sRow, " ")

	var result []int

	for _, elem := range arrSRow {
		if len(elem) == 0 {
			continue
		}
		num, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, num)
	}
	return result
}
