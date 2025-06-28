package main

import (
	"fmt"
	"os"
	"strconv"

	"hw7/internal/chessboard"
)

const usage string = "hw7 usage: hw7 [desk_size integer]\nExample: hw7 8"

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println(usage)
		return
	}

	size, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println(usage)
		return
	}

	chessboard.Draw(size)
}
