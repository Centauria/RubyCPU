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
	transTable map[[16]byte]transTableEntry
	running    bool
	result     chan *SearchResult
	controller chan struct{}
	mu         sync.Mutex
}

func NewEngine() *Ruby {
	return &Ruby{
		transTable: make(map[[16]byte]transTableEntry, Options["Hash"].(int)),
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
	r.transTable = make(map[[16]byte]transTableEntry, Options["Hash"].(int))
}

func (r *Ruby) WriteToMemory(pos *chess.Position, depth int8, typeFlag Flag, score float32) {
	r.transTable[pos.Hash()] = transTableEntry{
		depth:    depth,
		typeFlag: typeFlag,
		score:    score,
	}
}

func (r *Ruby) MiniMax(pos *chess.Position, depth int, isMaximizing bool, alpha, beta float32) {
	hash := pos.Hash()
	if _, ok := r.transTable[hash]; ok {

	} else {
		r.transTable[hash] = transTableEntry{
			depth:    0,
			typeFlag: 0,
			score:    0,
		}
	}
	if depth == 0 {

	}
}

func (r *Ruby) HandleSearch(_ Condition, timeout time.Duration) {
	r.running = true
	defer func() { r.running = false }()
	go func() {

	}()
	select {
	case <-time.After(timeout):
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
