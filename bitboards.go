package main

import (
	"fmt"
)

var bitTable = [64]int{
	63, 30, 3, 32, 25, 41, 22, 33, 15, 50, 42, 13, 11, 53, 19, 34, 61, 29, 2,
	51, 21, 43, 45, 10, 18, 47, 1, 54, 9, 57, 0, 35, 62, 31, 40, 4, 49, 5, 52,
	26, 60, 6, 23, 44, 46, 27, 56, 16, 7, 39, 48, 24, 59, 14, 12, 55, 38, 28,
	58, 20, 37, 17, 36, 8,
}

func PopBit(bb *uint64) int {
	var b uint64 = *bb ^ (*bb - 1)
	var fold uint32 = uint32(b&0xffffffff) ^ uint32(b>>32)

	*bb &= (*bb - 1)

	return bitTable[(fold*0x783a9b23)>>26]
}

func CountBits(b uint64) int {
	r := 0
	for b != 0 {
		b &= b - 1
		r++
	}

	return r
}

func ClearBit(bb *uint64, sq int) {
	*bb &= ClearMask[sq]
}

func SetBit(bb *uint64, sq int) {
	*bb |= SetMask[sq]
}

func PrintBitboard(bb uint64) {
	var shiftMe uint64 = 1

	rank := RANK_8
	file := FILE_A

	fmt.Println()

	for rank >= RANK_1 {
		for file <= FILE_H {
			sq := FR2SQ(file, rank)
			sq64 := SQ64(sq)

			if (shiftMe<<uint64(sq64))&bb != 0 {
				fmt.Printf("X")
			} else {
				fmt.Printf("-")
			}
			file++
		}
		fmt.Println()

		file = FILE_A
		rank--
	}

	fmt.Println()

}
