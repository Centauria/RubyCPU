package engine

import (
	"context"
	"github.com/notnil/chess"
	"time"
)

type Engine struct {
	Hash        int
	Threads     int
	timeManager TimeManager
	done        <-chan struct{}
	threads     []thread
	progress    func(info SearchInfo)
	mainLine    mainLine
	start       time.Time
	nodes       int64
	depth       int32
}

type thread struct {
	engine *Engine
	nodes  int64
	depth  int32
	stack  [stackSize]struct {
		position   *chess.Position
		moveList   []chess.Move
		pv         pv
		staticEval int
	}
}

type pv struct {
	items [stackSize]chess.Move
	size  int
}

type mainLine struct {
	moves []chess.Move
	score int
	depth int
}

type TimeManager interface {
	OnNodesChanged(nodes int)
	OnIterationComplete(line mainLine)
	Close()
}

func NewEngine() *Engine {
	return &Engine{
		Hash:    16,
		Threads: 1,
	}
}

func (e *Engine) Prepare() {
	if len(e.threads) != e.Threads {
		e.threads = make([]thread, e.Threads)
		for i := range e.threads {
			var t = &e.threads[i]
			t.engine = e
		}
	}
}

func (e *Engine) Search(ctx context.Context, searchParams SearchParams) SearchInfo {
	e.start = time.Now()
	e.Prepare()
	var p = &searchParams.Positions[len(searchParams.Positions)-1]
	ctx, e.timeManager = newSimpleTimeManager(ctx, e.start, searchParams.Limits, p)
	defer e.timeManager.Close()
	e.nodes = 0
	for i := range e.threads {
		var t = &e.threads[i]
		t.nodes = 0
		t.stack[0].position = p
	}
	e.progress = searchParams.Progress
	lazySmp(ctx, e)
	return e.currentSearchResult()
}

func (e *Engine) Clear() {
}
