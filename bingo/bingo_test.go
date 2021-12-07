package bingo

import (
	"fmt"
	"testing"
)

func TestIsWinner(t *testing.T) {
	var bb BingoBoard
	bb = append(bb, []int{1, 2, 3, 4, 5})
	bb = append(bb, []int{11, 22, 33, 44, 55})
	bb = append(bb, []int{12, 23, 34, 45, 56})
	bb = append(bb, []int{13, 24, 35, 46, 57})
	bb = append(bb, []int{14, 25, 36, 47, 58})

	got := bb.isWinner()
	fmt.Println("**********************************", got)
	if got == true {
		t.Errorf("winner = %t; want false", got)
	}

	bb = BingoBoard{}
	bb = append(bb, []int{1, 2, 3, 4, 5})
	bb = append(bb, []int{-1, -1, -1, -1, -1})
	bb = append(bb, []int{12, 23, 34, 45, 56})
	bb = append(bb, []int{13, 24, 35, 46, 57})
	bb = append(bb, []int{14, 25, 36, 47, 58})
	got = bb.isWinner()
	fmt.Println("**********************************", got)
	if got == false {
		t.Errorf("winner = %t; want false", got)
	}

	bb = BingoBoard{}
	bb = append(bb, []int{-1, 2, 3, 4, 5})
	bb = append(bb, []int{-1, 12, 13, 14, 15})
	bb = append(bb, []int{-1, 23, 34, 45, 56})
	bb = append(bb, []int{-1, 24, 35, 46, 57})
	bb = append(bb, []int{99, 25, 36, 47, 58})
	got = bb.isWinner()
	fmt.Println("**********************************", got)
	if got == true {
		t.Errorf("winner = %t; want false", got)
	}

	bb = BingoBoard{}
	bb = append(bb, []int{-1, 2, 3, -1, 5})
	bb = append(bb, []int{-1, 12, 13, -1, 15})
	bb = append(bb, []int{-1, 23, 34, -1, 56})
	bb = append(bb, []int{-1, 24, 35, -1, 57})
	bb = append(bb, []int{99, 25, 36, -1, 58})
	got = bb.isWinner()
	fmt.Println("**********************************", got)
	if got == false {
		t.Errorf("winner = %t; want false", got)
	}

	bb = BingoBoard{}
	bb = append(bb, []int{-1, 2, 3, 4, 5})
	bb = append(bb, []int{-1, -1, 13, -1, 15})
	bb = append(bb, []int{-1, 23, 34, -1, -1})
	bb = append(bb, []int{-1, 24, -1, -1, 57})
	bb = append(bb, []int{99, 25, 36, -1, 58})
	got = bb.isWinner()
	fmt.Println("**********************************", got)
	if got == true {
		t.Errorf("winner = %t; want false", got)
	}

}
