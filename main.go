package main

import (
	"fmt"
)

var Sq120ToSq64 [BRD_SQ_NUM]int
var Sq64Tosq120 [64]int

func initSq120To64() {
	file := FILE_A
	rank := RANK_1
	sq64 := 0

	for i := range BRD_SQ_NUM {
		Sq120ToSq64[i] = 65
	}

	for i := range 64 {
		Sq64Tosq120[i] = 120
	}

	for rank <= RANK_8 {
		for file <= FILE_H {
			sq := FR2SQ(file, rank)
			Sq64Tosq120[sq64] = sq
			Sq120ToSq64[sq] = sq64
			sq64++
			file++
		}
		file = FILE_A
		rank++
	}

}

func initAll() {
	initSq120To64()
}

func printBoard() {
	for i, v := range Sq120ToSq64 {
		if i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d", v)
	}

	fmt.Println()
	fmt.Println()

	for i, v := range Sq64Tosq120 {
		if i%8 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d", v)
	}
	fmt.Println()
}

func main() {
	fmt.Println("Hello, Chess!")

	initAll()

	// printBoard()
	num1 := 1
	num2 := 2

	assert(num1 != num2, "Debe ser iguales")
}
