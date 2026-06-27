package main

import (
	"errors"
	"fmt"
)

func ResetBoard(pos *SBoard) {
	for i := range BRD_SQ_NUM {
		pos.Pieces[i] = OFFBOARD
	}

	for i := range 64 {
		pos.Pieces[SQ120(i)] = EMPTY
	}

	for i := range 3 {
		pos.BigPce[i] = 0
		pos.MajPce[i] = 0
		pos.MinPce[i] = 0
		pos.Pawns[i] = uint64(0)
	}

	for i := range 13 {
		pos.PceNums[i] = 0
	}

	pos.KingsSqr[WHITE], pos.KingsSqr[BLACK] = NO_SQ, NO_SQ

	pos.Side = BOTH
	pos.EnPass = NO_SQ
	pos.FiftyMove = 0

	pos.Ply = 0
	pos.HisPly = 0

	pos.CastlePerm = 0

	pos.PosKey = uint64(0)
}

func PrintBoard(pos *SBoard){
	
}

func ParseFEN(fen string, pos *SBoard) error {
	assert(fen != "", "Debe haber un valor FEN")
	assert(pos != nil, "Debe haber un tablero")

	rank := RANK_8
	file := FILE_A
	piece := 0
	count := 0
	index := 0

	ResetBoard(pos)

	for i, c := range fen {
		if rank < 0 {
			index = i
			break
		}
		count = 1
		switch c {
		case 'p':
			piece = BP
		case 'r':
			piece = BR
		case 'n':
			piece = BN
		case 'b':
			piece = BB
		case 'q':
			piece = BQ
		case 'k':
			piece = BK
		case 'P':
			piece = WP
		case 'R':
			piece = WR
		case 'N':
			piece = WN
		case 'B':
			piece = WB
		case 'Q':
			piece = WQ
		case 'K':
			piece = WK

		case '1', '2', '3', '4', '5', '6', '7', '8':
			piece = EMPTY
			count = int(c - '0')

		case '/', ' ':
			rank--
			file = FILE_A
			continue

		default:
			fmt.Printf("FEN ERROR \n")
			return errors.New("FEN ERROR")
		}

		for i := 0; i < count; i++ {
			sq64 := rank*8 + file
			sq120 := SQ120(sq64)
			if piece != EMPTY {
				pos.Pieces[sq120] = piece
			}
			file++
		}

	}

	assert(fen[index] == 'w' || fen[index] == 'b', "Debe indicarse si es negras 'b' o blancas 'w'")

	if fen[index] == 'w' {
		pos.Side = WHITE
	} else {
		pos.Side = BLACK
	}

	index += 2

	for i := 0; i < 4; i++ {
		if fen[index] == ' ' {
			break
		}
		switch fen[index] {
		case 'K':
			pos.CastlePerm |= WKCA
		case 'Q':
			pos.CastlePerm |= WQCA
		case 'k':
			pos.CastlePerm |= BKCA
		case 'q':
			pos.CastlePerm |= BQCA
		}
		index++
	}
	index++

	assert(pos.CastlePerm >= 0 && pos.CastlePerm <= 15, "Castle Perm debe ser mayor a cero y menor a 16")

	if fen[index] != '-' {
		file = int(fen[index] - 'a')
		rank = int(fen[index+1] - '1')

		assert(file >= FILE_A && file <= FILE_H, "El valor de la fila debe estar entre FILE_A y FILE_H")
		assert(rank >= RANK_1 && rank <= RANK_8, "El valor del rango debe estar entre RANK_1 y RANK_8")

		pos.EnPass = FR2SQ(file, rank)

	}

	pos.PosKey = GeneratePosKey(pos)

	return nil
}
