package engine

import "github.com/notnil/chess"

// Options available:
// Hash: int
// UCI_AnalyseMode: bool
var (
	Options = map[string]interface{}{
		"Hash":            1,
		"UCI_AnalyseMode": false,
		"UCI_Chess960":    false,
	}
	//TODO: add `OptionConditions` which includes limits of the Options
)

type Engine interface {
	Search(receiver chan string, condition Condition)
	Evaluate(position *chess.Position)
}

type Ruby struct {
	game *chess.Game
	pv   []*chess.Move
}

func NewEngine() *Ruby {
	return &Ruby{}
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

func (r *Ruby) Search(ch chan string, condition Condition) {

}

func (r *Ruby) Evaluate(position *chess.Position) {

}
