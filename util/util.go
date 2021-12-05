package util

import (
	"fmt"
	"os"
	"strconv"
)

func GetIntFromString(s string) int {
	iCur, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return iCur
}
