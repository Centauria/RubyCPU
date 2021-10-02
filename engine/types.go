package engine

import (
	"github.com/notnil/chess"
	"time"
)

type LimitsType struct {
	Ponder         bool
	Infinite       bool
	WhiteTime      int
	BlackTime      int
	WhiteIncrement int
	BlackIncrement int
	MoveTime       int
	MovesToGo      int
	Depth          int
	Nodes          int
	Mate           int
}

type SearchParams struct {
	Positions []chess.Position
	Limits    LimitsType
	Progress  func(si SearchInfo)
}

type SearchInfo struct {
	Score    UciScore
	Depth    int
	Nodes    int64
	Time     time.Duration
	MainLine []chess.Move
}

type UciScore struct {
	Centipawns int
	Mate       int
}
