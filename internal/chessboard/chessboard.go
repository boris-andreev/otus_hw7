package chessboard

import "fmt"

func Draw(size int) {
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			fmt.Print(getCallContent(x + y))
		}
		fmt.Print("\n")
	}
}

func getCallContent(cellId int) string {
	if cellId%2 == 0 {
		return "X"
	}

	return " "
}
