package main

const BRD_SQ_NUM = 120
const NAME = "JChess v1.0"
const MAX_GAME_MOVES = 2048

// Pieces List
const (
	Empty = iota
	WP    // White Pawn
	WN    // White Knight
	WB    // White Bishop
	WR    // White Rook
	WQ    // White Queen
	WK    // White King
	BP    // Black Pawn
	BN    // Black Knight
	BB    // Black Bishop
	BR    // Black Rook
	BQ    // Black Queen
	BK    // Black King
)

// Board Files
const (
	FILE_A = iota
	FILE_B
	FILE_C
	FILE_D
	FILE_E
	FILE_F
	FILE_G
	FILE_H
	FILE_NONE
)

// Board Ranks
const (
	RANK_1 = iota
	RANK_2
	RANK_3
	RANK_4
	RANK_5
	RANK_6
	RANK_7
	RANK_8
	RANK_EMPTY
)

// Players
const (
	WHITE = iota
	BLACK
	BOTH
)

const (
	WKCA = 1
	WQCA = 2
	BKCA = 4
	BQCA = 8
)

// Board Square
const (
	A1 = iota + 21
	B1
	C1
	D1
	E1
	F1
	G1
	H1
)

const (
	A2 = iota + 31
	B2
	C2
	D2
	E2
	F2
	G2
	H2
)

const (
	A3 = iota + 41
	B3
	C3
	D3
	E3
	F3
	G3
	H3
)

const (
	A4 = iota + 51
	B4
	C4
	D4
	E4
	F4
	G4
	H4
)

const (
	A5 = iota + 61
	B5
	C5
	D5
	E5
	F5
	G5
	H5
)

const (
	A6 = iota + 71
	B6
	C6
	D6
	E6
	F6
	G6
	H6
)

const (
	A7 = iota + 81
	B7
	C7
	D7
	E7
	F7
	G7
	H7
)

const (
	A8 = iota + 91
	B8
	C8
	D8
	E8
	F8
	G8
	H8
	NO_SQ
)

// Board Configuration
type SUndo struct {
	Move       int
	CastlePerm int
	EnPass     int
	FiftyMove  int
	PosKey     uint64
}

type SBoard struct {
	Pieces [BRD_SQ_NUM]int
	Pawns  [3]uint64

	KingsSqr [2]int

	Side      int
	EnPass    int
	FiftyMove int

	CastlePerm int

	Ply    int
	HisPly int

	PosKey  uint64
	PceNums [13]int

	BigPce [3]int
	MajPce [3]int
	MinPce [3]int

	History [MAX_GAME_MOVES]SUndo

	Plist [13][10]int
}

// var Sq120ToSq64 [BRD_SQ_NUM]int
// var Sq64Tosq120 [64]int

func FR2SQ(f, r int) int {
	return (21 + f) + (r * 10)
}
