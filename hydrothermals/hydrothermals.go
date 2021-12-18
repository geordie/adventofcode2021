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

func ParseInputPuzzle1(sFile string) Grid {
	file, err := os.Open(sFile)
	if err != nil {
		log.Fatal(err)
	}

	maxX := 0
	maxY := 0
	lines := []Line{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := parseLine(scanner.Text())
		if !line.IsHorizontal() && !line.IsVertical() {
			continue
		}
		lines = append(lines, line)
		maxX = util.Max(maxX, util.Max(line.x1, line.x2)) + 1
		maxY = util.Max(maxY, util.Max(line.y1, line.y2)) + 1
	}

	grid := buildGrid(maxX, maxY)
	for _, line := range lines {
		grid.AddAnyLine(line)
	}
	return grid
}

func ParseInputPuzzle2(sFile string) Grid {
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
		lines = append(lines, line)
		maxX = util.Max(maxX, util.Max(line.x1, line.x2)) + 1
		maxY = util.Max(maxY, util.Max(line.y1, line.y2)) + 1
	}

	grid := buildGrid(maxX, maxY)

	for _, line := range lines {
		grid.AddAnyLine(line)
	}

	return grid
}

func buildGrid(x int, y int) Grid {

	iGrid := make([][]int, x)
	for i := range iGrid {
		iGrid[i] = make([]int, y)
	}

	grid := Grid{
		table: iGrid,
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

func (line Line) IsSlopePositive() bool {
	xDiff := line.x2 - line.x1
	yDiff := line.y2 - line.y1

	// Avoid dividing by zero
	if yDiff == 0 {
		return true
	}
	return (xDiff / yDiff) > 0
}

func (line Line) IsSlopeNegative() bool {
	return !line.IsSlopePositive()
}

func (grid Grid) AddAnyLine(line Line) {
	minX := util.Min(line.x1, line.x2)
	maxX := util.Max(line.x1, line.x2)
	minY := util.Min(line.y1, line.y2)
	maxY := util.Max(line.y1, line.y2)

	if line.IsHorizontal() {
		for i := minY; i <= maxY; i++ {
			grid.table[line.x1][i]++
		}
	} else if line.IsVertical() {
		for i := minX; i <= maxX; i++ {
			grid.table[i][line.y1]++
		}
	} else if line.IsSlopePositive() {
		for i := 0; i <= maxX-minX; i++ {
			grid.table[minX+i][minY+i]++
		}
	} else if line.IsSlopeNegative() {
		for i := 0; i <= maxX-minX; i++ {
			grid.table[minX+i][maxY-i]++
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
