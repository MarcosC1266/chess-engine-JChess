package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Chess!")

	InitAll()

	board := SBoard{}

	ResetBoard(&board)
	ParseFEN(START_FEN, &board)

}
