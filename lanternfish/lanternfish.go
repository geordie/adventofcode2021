package lanternfish

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/geordie/adventofcode2021/util"
)

type FishList []int
type FishHash [9]int

func ParseInput(sFile string) FishList {
	file, err := os.Open(sFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	sFishList := scanner.Text()
	fishList := parseFishList(sFishList)
	return fishList
}

func ParseInput2(sFile string) FishHash {
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

// Build a slice of ints from comma delimited list of strings
func parseFishList(sFishList string) FishList {
	stringFishList := strings.Split(sFishList, ",")
	fishList := make([]int, len(stringFishList))
	for i, elem := range stringFishList {
		fishList[i] = util.GetIntFromString(elem)
	}
	return fishList
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

func (fishList FishList) IterateModel(iterations int) FishList {
	for i := 0; i < iterations; i++ {
		fishList = fishList.iterateOnce()
	}
	return fishList
}

func (fishList FishList) iterateOnce() FishList {
	fishListCopy := make([]int, len(fishList))
	copy(fishListCopy, fishList)

	for i, elem := range fishList {
		if elem > 0 {
			fishListCopy[i]--
		} else if elem == 0 {
			fishListCopy[i] = 6
			fishListCopy = append(fishListCopy, 8)
		}
	}
	return fishListCopy
}
