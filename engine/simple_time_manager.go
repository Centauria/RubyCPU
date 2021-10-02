package engine

import (
	"context"
	"github.com/notnil/chess"
	"time"
)

type simpleTimeManager struct {
	start     time.Time
	limits    LimitsType
	softLimit time.Duration
	hardLimit time.Duration
	cancel    context.CancelFunc
}

func newSimpleTimeManager(ctx context.Context, start time.Time,
	limits LimitsType, p *chess.Position) (context.Context, *simpleTimeManager) {

	var tm = &simpleTimeManager{
		start:  start,
		limits: limits,
	}

	if limits.MoveTime > 0 {
		tm.hardLimit = time.Duration(limits.MoveTime) * time.Millisecond
	} else if limits.WhiteTime > 0 || limits.BlackTime > 0 {
		var main, inc time.Duration
		if p.Turn() == chess.White {
			main = time.Duration(limits.WhiteTime) * time.Millisecond
			inc = time.Duration(limits.WhiteIncrement) * time.Millisecond
		} else {
			main = time.Duration(limits.BlackTime) * time.Millisecond
			inc = time.Duration(limits.BlackIncrement) * time.Millisecond
		}
		tm.softLimit, tm.hardLimit = calcLimits(main, inc, limits.MovesToGo)
	}

	var cancel context.CancelFunc
	if tm.hardLimit != 0 {
		ctx, cancel = context.WithDeadline(ctx, start.Add(tm.hardLimit))
	} else {
		ctx, cancel = context.WithCancel(ctx)
	}

	tm.cancel = cancel
	return ctx, tm
}

func (tm *simpleTimeManager) OnNodesChanged(nodes int) {
	if tm.limits.Nodes > 0 && nodes >= tm.limits.Nodes {
		tm.cancel()
	}
}

func (tm *simpleTimeManager) OnIterationComplete(line mainLine) {
	if tm.limits.Infinite {
		return
	}
	if tm.limits.Depth != 0 && line.depth >= tm.limits.Depth {
		tm.cancel()
		return
	}
	if line.score >= winIn(line.depth-5) ||
		line.score <= lossIn(line.depth-5) {
		tm.cancel()
		return
	}
	if tm.softLimit != 0 &&
		time.Since(tm.start) >= tm.softLimit {
		tm.cancel()
		return
	}
}

func (tm *simpleTimeManager) Close() {
	tm.cancel()
}

func calcLimits(main, inc time.Duration, moves int) (soft, hard time.Duration) {
	const (
		DefaultMovesToGo = 40
		MoveOverhead     = 300 * time.Millisecond
		MinTimeLimit     = 1 * time.Millisecond
	)

	main -= MoveOverhead
	if main < MinTimeLimit {
		main = MinTimeLimit
	}

	if moves == 0 {
		hard = main/13 + inc
		soft = hard / 4
	} else {
		if DefaultMovesToGo < moves {
			moves = DefaultMovesToGo
		}
		hard = main/time.Duration(moves) + inc
		soft = hard / 2
	}

	hard = limitDuration(hard, MinTimeLimit, main)
	soft = limitDuration(soft, MinTimeLimit, main)

	return
}

func limitDuration(v, min, max time.Duration) time.Duration {
	if v < min {
		return min
	}
	if v > max {
		return max
	}
	return v
}
