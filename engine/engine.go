package engine

import (
	"fmt"
	"github.com/Centauria/RubyCPU/eval"
	"github.com/notnil/chess"
	"strings"
	"sync"
	"time"
)

type Engine interface {
	Search(condition Condition)
	Evaluate(position *chess.Position) float32
	Stop()
}

type Ruby struct {
	game       *chess.Game
	pv         []*chess.Move
	transTable []transTableEntry
	running    bool
	result     chan *SearchResult
	controller chan struct{}
	mu         sync.Mutex
}

func NewEngine() *Ruby {
	return &Ruby{
		transTable: make([]transTableEntry, Options["Hash"].(int)),
		result:     make(chan *SearchResult),
		controller: make(chan struct{}),
	}
}

func (r *Ruby) NewGame() {
	r.game = chess.NewGame(chess.UseNotation(chess.UCINotation{}))
}

func (r *Ruby) SetPosition(fen string) (err error) {
	f, err := chess.FEN(fen)
	if err == nil {
		f(r.game)
	}
	return
}

func (r *Ruby) Move(moveStr string) (err error) {
	err = r.game.MoveStr(moveStr)
	return
}

func (r *Ruby) ShowBoard() {
	println(r.game.Position().Board().Draw())
}

func (r *Ruby) PositionString() string {
	moves := make([]string, len(r.game.Moves()))
	for i, move := range r.game.Moves() {
		moves[i] = move.String()
	}
	return strings.Join(moves, " ")
}

func (r *Ruby) ResetMemory() {
	for _, tt := range r.transTable {
		tt.key = 0
		tt.score = 0
		tt.depth = 0
		tt.typeFlag = Exact
	}
}

func (r *Ruby) HandleSearch(_ Condition, timeout time.Duration) {
	r.running = true
	defer func() { r.running = false }()
	select {
	case <-time.After(timeout):
		r.result <- &SearchResult{
			BestMove: "d2d4",
			Ponder:   "d7d5",
		}
		return
	case <-r.controller:
		r.result <- &SearchResult{
			BestMove: "e2e4",
			Ponder:   "e7e5",
		}
		return
	}
}

func (r *Ruby) ShowResult() {
	var result *SearchResult
	for {
		select {
		case result = <-r.result:
			fmt.Printf("bestmove %s ponder %s\n", result.BestMove, result.Ponder)
			return
		}
	}
}

func (r *Ruby) Search(condition Condition) {
	go r.HandleSearch(condition, 10*time.Second)
	go r.ShowResult()
}

func (r *Ruby) Evaluate(position *chess.Position) float32 {
	mb := eval.MaterialImbalance(position)
	return mb
}

func (r *Ruby) Stop() {
	if r.running {
		r.controller <- struct{}{}
	}
}
