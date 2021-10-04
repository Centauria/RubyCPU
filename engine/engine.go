package engine

import "github.com/notnil/chess"

// Options available:
// Hash: int
// UCI_AnalyseMode: bool
var (
	Options = map[string]interface{}{
		"Hash":            1,
		"UCI_AnalyseMode": false,
	}
)

type Engine struct {
	game *chess.Game
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) NewGame() {
	e.game = chess.NewGame(chess.UseNotation(chess.UCINotation{}))
}

func (e *Engine) SetPosition(fen string) (err error) {
	f, err := chess.FEN(fen)
	if err == nil {
		f(e.game)
	}
	return
}

func (e *Engine) Move(moveStr string) (err error) {
	err = e.game.MoveStr(moveStr)
	return
}

func (e *Engine) ShowBoard() {
	println(e.game.Position().Board().Draw())
}
