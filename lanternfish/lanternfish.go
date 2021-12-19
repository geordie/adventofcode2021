package lanternfish

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/geordie/adventofcode2021/util"
)

type FishList []int

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

// Build a slice of ints from comma delimited list of strings
func parseFishList(sFishList string) FishList {
	stringFishList := strings.Split(sFishList, ",")
	fishList := make([]int, len(stringFishList))
	for i, elem := range stringFishList {
		fishList[i] = util.GetIntFromString(elem)
	}
	return fishList
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
