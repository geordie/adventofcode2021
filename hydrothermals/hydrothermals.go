package hydrothermals

import (
	"bufio"
	"log"
	"os"
	"strings"

	util "github.com/geordie/adventofcode2021/util"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

type Grid struct {
	table [][]int
}

func ParseInput(sFile string) Grid {
	file, err := os.Open(sFile)
	if err != nil {
		log.Fatal(err)
	}

	maxX := 0
	maxY := 0
	lines := []Line{}
	scanner := bufio.NewScanner(file)

	// Get the BingoNumbers for first row of file
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		if !line.IsHorizontal() && !line.IsVertical() {
			continue
		}
		lines = append(lines, line)
		maxX = util.Max(maxX, util.Max(line.x1, line.x2)) + 1
		maxY = util.Max(maxY, util.Max(line.y1, line.y2)) + 1
	}

	iGrid := make([][]int, maxX)
	for i := range iGrid {
		iGrid[i] = make([]int, maxY)
	}

	grid := Grid{
		table: iGrid,
	}

	for _, line := range lines {
		grid.AddLine(line)
	}

	return grid
}

func parseLine(sLine string) Line {
	sPoints := strings.Fields(sLine)
	sCoordinates := strings.Split(sPoints[0]+","+sPoints[2], ",")

	line := Line{
		x1: util.GetIntFromString(sCoordinates[0]),
		y1: util.GetIntFromString(sCoordinates[1]),
		x2: util.GetIntFromString(sCoordinates[2]),
		y2: util.GetIntFromString(sCoordinates[3]),
	}

	if line.x1 > line.x2 {
		temp := line.x1
		line.x1 = line.x2
		line.x2 = temp
	}

	if line.y1 > line.y2 {
		temp := line.y1
		line.y1 = line.y2
		line.y2 = temp
	}

	return line
}

func (line Line) IsHorizontal() bool {
	if line.x1 == line.x2 {
		return true
	}
	return false
}

func (line Line) IsVertical() bool {
	if line.y1 == line.y2 {
		return true
	}
	return false
}

func (grid Grid) AddLine(line Line) {

	if line.IsHorizontal() {
		for i := line.y1; i <= line.y2; i++ {
			grid.table[line.x1][i]++
		}
	} else if line.IsVertical() {
		for i := line.x1; i <= line.x2; i++ {
			grid.table[i][line.y1]++
		}
	}
}

func (grid Grid) CountNoGos() int {
	result := 0
	for i := range grid.table {
		for j := range grid.table[i] {
			if grid.table[i][j] > 1 {
				result++
			}
		}
	}
	return result
}
