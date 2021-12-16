package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	bingo "github.com/geordie/adventofcode2021/bingo"
	diag "github.com/geordie/adventofcode2021/diagnostics"
	hydro "github.com/geordie/adventofcode2021/hydrothermals"
	nav "github.com/geordie/adventofcode2021/navigation"
	util "github.com/geordie/adventofcode2021/util"
)

func main() {
	day5puzzle1()
	day4puzzle2()
	day4puzzle1()
	day3puzzle2()
	day3puzzle1()
	day2puzzle2()
	day2puzzle1()
	day1puzzle1()
	day1puzzle2()
}

func day5puzzle1() {
	grid := hydro.ParseInput("input/day5.txt")
	result := grid.CountNoGos()
	fmt.Println("ANSWER: ", result)

}

func day4puzzle2() {
	bingoNums, bingoBoards := bingo.ParseInput("input/day4.txt")
	_, score := bingoBoards.RunGameToLose(bingoNums)
	fmt.Println("DAY 4, ANSWER 2:", score)
}

func day4puzzle1() {
	bingoNums, bingoBoards := bingo.ParseInput("input/day4.txt")
	_, score := bingoBoards.RunGame(bingoNums)
	fmt.Println("DAY 4, ANSWER 1:", score)
}

func day3puzzle2() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}

	var diagnostics diag.Diagnostics
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		d := diag.ParseDiagnostic(s)
		diagnostics = append(diagnostics, d)
	}

	dArrO2 := diagnostics.Filter(0, diag.O2)
	dO2 := dArrO2[0]
	iO2 := dO2.Int()

	dArrCO2 := diagnostics.Filter(0, diag.CO2)
	dCO2 := dArrCO2[0]
	iCO2 := dCO2.Int()

	fmt.Println("DAY 3, ANSWER 2:", iO2*iCO2)
}

func day3puzzle1() {
	file, err := os.Open("input/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	var mainD diag.Diagnostic

	total := diag.Counter(0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		total++
		d := diag.ParseDiagnostic(s)
		mainD.Add(d)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Build Gamma
	dGamma := mainD.Gamma(total)
	iGamma := dGamma.Int()

	// Build Epsilon
	dEpsilon := mainD.Epsilon(total)
	iEpsilon := dEpsilon.Int()

	// Answer is product of Gamma and Epsilon
	iAnswer := iGamma * iEpsilon

	fmt.Println("DAY 3, ANSWER 1:", iAnswer)
}

func day2puzzle2() {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	location := nav.Position{
		Aim: 0,
		Forward: nav.Vector{
			Direction: nav.Forward,
			Magnitude: 0,
		},
		Down: nav.Vector{
			Direction: nav.Down,
			Magnitude: 0},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		v := nav.ParseDirection(s)
		location.Move(v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DAY 2, ANSWER 2:", location.Product())
}

func day2puzzle1() {
	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	location := nav.Position{
		Forward: nav.Vector{
			Direction: nav.Forward,
			Magnitude: 0,
		},
		Down: nav.Vector{
			Direction: nav.Down,
			Magnitude: 0},
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		v := nav.ParseDirection(s)
		location.Add(v)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DAY 2, ANSWER 1:", location.Product())
}

func day1puzzle2() {

	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	i1 := util.GetIntFromString(scanner.Text())
	scanner.Scan()
	i2 := util.GetIntFromString(scanner.Text())
	scanner.Scan()
	i3 := util.GetIntFromString(scanner.Text())

	iLastWindowSum := i1 + i2 + i3
	iCountOfIncreases := 0

	for scanner.Scan() {
		s := scanner.Text()
		i4 := util.GetIntFromString(s)
		iCurWindowSum := i2 + i3 + i4
		if iCurWindowSum > iLastWindowSum {
			iCountOfIncreases++
		}
		iLastWindowSum = iCurWindowSum
		i2 = i3
		i3 = i4
	}

	fmt.Println("DAY 1, ANSWER 2:", iCountOfIncreases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day1puzzle1() {

	file, err := os.Open("input/day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	iLast := 0
	iCountOfIncreases := -1

	for scanner.Scan() {
		s := scanner.Text()
		iCur := util.GetIntFromString(s)
		if iCur > iLast {
			iCountOfIncreases++
		}
		iLast = iCur
	}

	fmt.Println("DAY 1, ANSWER 1:", iCountOfIncreases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
