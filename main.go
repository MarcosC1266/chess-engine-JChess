package main

import (
	"chess/config"
	"fmt"
)

var Sq120ToSq64 [config.BRD_SQ_NUM]int
var Sq64Tosq120 [64]int

func initSq120To64() {
	file := config.FILE_A
	rank := config.RANK_1
	sq64 := 0

	for i := range config.BRD_SQ_NUM {
		Sq120ToSq64[i] = 65
	}

	for i := range 64 {
		Sq64Tosq120[i] = 120
	}

	for rank <= config.RANK_8 {
		for file <= config.FILE_H {
			sq := config.FR2SQ(file, rank)
			Sq64Tosq120[sq64] = sq
			Sq120ToSq64[sq] = sq64
			sq64++
			file++
		}
		file = config.FILE_A
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

	printBoard()
}
