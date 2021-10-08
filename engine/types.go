package engine

import (
	"time"
)

type Condition struct {
	Depth     int
	Nodes     uint64
	MoveTime  int // in milliseconds
	Infinite  bool
	StartTime time.Time
	LastTime  time.Time
}

type UciScore struct {
	Centipawns int
	Mate       int
}

type SearchResult struct {
	BestMove string
	Ponder   string
}
