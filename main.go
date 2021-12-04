package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

	nav "github.com/geordie/adventofcode2021/navigation"
)

func main() {
	day2puzzle2()
	day2puzzle1()
	//day1puzzle1()
	//day1puzzle2()
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

	file, err := os.Open("input/day2.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	i1 := getIntFromString(scanner.Text())
	scanner.Scan()
	i2 := getIntFromString(scanner.Text())
	scanner.Scan()
	i3 := getIntFromString(scanner.Text())

	iLastWindowSum := i1 + i2 + i3
	iCountOfIncreases := 0

	for scanner.Scan() {
		s := scanner.Text()
		i4 := getIntFromString(s)
		iCurWindowSum := i2 + i3 + i4
		if iCurWindowSum > iLastWindowSum {
			iCountOfIncreases++
		}
		iLastWindowSum = iCurWindowSum
		i2 = i3
		i3 = i4
	}

	fmt.Println("ANSWER #2:", iCountOfIncreases)

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
		iCur := getIntFromString(s)
		if iCur > iLast {
			iCountOfIncreases++
		}
		iLast = iCur
	}

	fmt.Println("ANSWER #1:", iCountOfIncreases)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getIntFromString(s string) int {
	iCur, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return iCur
}
