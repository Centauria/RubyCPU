package eval

import "github.com/notnil/chess"

func MaterialImbalance(position *chess.Position) float32 {
	board := position.Board()
	var materials [2]float32
	for sq := 0; sq < 64; sq++ {
		p := board.Piece(chess.Square(sq))
		color := -1
		switch p.Color() {
		case chess.White:
			color = 0
		case chess.Black:
			color = 1
		}
		if color != -1 {
			switch p.Type() {
			case chess.Pawn:
				materials[color] += 1.0
			case chess.Knight:
				materials[color] += 3.0
			case chess.Bishop:
				materials[color] += 3.0
			case chess.Rook:
				materials[color] += 5.0
			case chess.Queen:
				materials[color] += 9.0
			}
		}
	}
	return materials[0] - materials[1]
}
