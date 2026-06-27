package main

func GeneratePosKey(pos *SBoard) uint64 {

	var finalKey uint64
	sq := 0
	piece := EMPTY

	// pieces
	for sq < BRD_SQ_NUM {
		piece = pos.Pieces[sq]
		if piece != NO_SQ && piece != EMPTY && piece != OFFBOARD {
			assert(piece >= WP && piece <= BK, "Debe ser una pieza valida")
			finalKey ^= PiecesKeys[piece][sq]
		}
		sq++
	}

	if pos.Side == WHITE {
		finalKey ^= SideKey
	}

	if pos.EnPass != NO_SQ {
		assert(pos.EnPass >= 0 && pos.EnPass < BRD_SQ_NUM, "Aun no se que significa esto")
		finalKey ^= PiecesKeys[EMPTY][pos.EnPass]
	}

	assert(pos.CastlePerm >= 0 && pos.CastlePerm <= 15, "El permiso de enroque debe estar entre 0 y 15")
	finalKey ^= CastleKeys[pos.CastlePerm]

	return finalKey
}
