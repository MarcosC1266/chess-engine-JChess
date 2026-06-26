package main

import (
	"math/rand"
)

var Sq120ToSq64 [BRD_SQ_NUM]int
var Sq64Tosq120 [64]int

var SetMask [64]uint64
var ClearMask [64]uint64

var PiecesKeys [13][120]uint64
var SideKey uint64
var CastleKeys [16]uint64

func InitAll() {
	initSq120To64()
	initBitMasks()
	initHashKeys()
}

func initHashKeys() {
	for i := range 13 {
		for j := range 120 {
			PiecesKeys[i][j] = rand.Uint64()
		}
	}

	SideKey = rand.Uint64()

	for i := range 16 {
		CastleKeys[i] = rand.Uint64()
	}
}

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
func initBitMasks() {
	for i := range 64 {
		SetMask[i] = uint64(0)
		ClearMask[i] = uint64(0)
	}

	for i := range 64 {
		SetMask[i] |= (uint64(1) << i)
		ClearMask[i] = ^SetMask[i]
	}
}
