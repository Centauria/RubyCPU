package engine

import "github.com/notnil/chess"

type Flag int8

const (
	Exact Flag = iota
	Alpha
	Beta
)

type transTableEntry struct {
	depth    int8
	typeFlag Flag
	score    float32
	bestMove *chess.Move
}

//`HashSize` defined as length of hash table temporarily

type TransPositionTree struct {
	pos      *transTableEntry
	children []*TransPositionTree
}
