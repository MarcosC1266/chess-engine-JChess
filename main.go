package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Chess!")

	InitAll()

	var playBitboard uint64 = 0

	SetBit(&playBitboard, SQ64(A8))
	SetBit(&playBitboard, SQ64(B8))
	SetBit(&playBitboard, SQ64(C8))
	SetBit(&playBitboard, SQ64(D8))
	SetBit(&playBitboard, SQ64(E8))
	SetBit(&playBitboard, SQ64(F8))
	SetBit(&playBitboard, SQ64(G8))
	SetBit(&playBitboard, SQ64(H8))

	PrintBitboard(playBitboard)

	ClearBit(&playBitboard, SQ64(E8))

	PrintBitboard(playBitboard)

}
