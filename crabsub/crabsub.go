package crabsub

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/geordie/adventofcode2021/util"
)

type CrabSubPoints []int

func ParseInput(sFile string) CrabSubPoints {
	file, err := os.Open(sFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	sCrabSubPoints := scanner.Text()
	crabSubPoints := parseList(sCrabSubPoints)
	return crabSubPoints
}

func parseList(sCrabSubPoints string) CrabSubPoints {
	stringsCrabSubPoints := strings.Split(sCrabSubPoints, ",")
	crabSubPoints := make([]int, len(stringsCrabSubPoints))
	for _, elem := range stringsCrabSubPoints {
		iValue := util.GetIntFromString(elem)
		crabSubPoints = append(crabSubPoints, iValue)
	}
	return crabSubPoints
}
