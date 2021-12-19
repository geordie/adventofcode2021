package lanternfish

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/geordie/adventofcode2021/util"
)

type FishHash [9]int

func ParseInput(sFile string) FishHash {
	file, err := os.Open(sFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	sFishList := scanner.Text()
	fishHash := parseFishList2(sFishList)
	return fishHash
}

func parseFishList2(sFishList string) FishHash {
	stringsFishList := strings.Split(sFishList, ",")
	fishHash := FishHash{}
	for _, elem := range stringsFishList {
		iValue := util.GetIntFromString(elem)
		fishHash[iValue]++
	}
	return fishHash
}

func (fishHash FishHash) IterateModel(iterations int) FishHash {
	for i := 0; i < iterations; i++ {
		fishHash = fishHash.iterateOnce()
	}
	return fishHash
}

func (fishHash FishHash) iterateOnce() FishHash {
	iNewFish := fishHash[0]
	for i := 0; i < 8; i++ {
		fishHash[i] = fishHash[i+1]
	}
	fishHash[8] = iNewFish
	fishHash[6] += iNewFish
	return fishHash
}

func (fishHash FishHash) TotalFish() int {
	iFish := 0
	for i := 0; i < 9; i++ {
		iFish += fishHash[i]
	}
	return iFish
}
