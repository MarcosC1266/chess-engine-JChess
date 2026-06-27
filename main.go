package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Chess!")

	InitAll()

	board := SBoard{}

	ParseFEN(START_FEN, &board)
	PrintBoard(&board)

	ParseFEN(FEN_1, &board)
	PrintBoard(&board)

	ParseFEN(FEN_2, &board)
	PrintBoard(&board)

	ParseFEN(FEN_3, &board)
	PrintBoard(&board)

}
