package uci

import (
	"context"
	"fmt"
	"github.com/Centauria/RubyCPU/engine"
	"github.com/notnil/chess"
	"strconv"
	"strings"
)

type Protocol interface {
	Prepare()
	Process()
	Cleanup()
}

type ProtocolUCI struct {
	brain *engine.Ruby
}

func (p *ProtocolUCI) Prepare() {
	p.brain = engine.NewEngine()
}

func (p *ProtocolUCI) Process() {

}

func (p *ProtocolUCI) Cleanup() {

}

func (p *ProtocolUCI) Handle(_ context.Context, command string) error {
	cmdArray := strings.Split(command, " ")
	switch cmdArray[0] {
	case "uci":
		println("id name Ruby")
		println("id author Centauria CHEN")
		println("option name Hash type spin default 1 min 1 max 128")
		println("uciok")
		//TODO: show all options using a function instead
	case "setoption":
		var name, value string
		for i := 1; i < len(cmdArray); i++ {
			if cmdArray[i] == "name" {
				i++
				name = cmdArray[i]
			} else if cmdArray[i] == "value" {
				i++
				value = cmdArray[i]
			}
		}
		err := p.setOption(name, value)
		if err != nil {
			return err
		}
	case "isready":
		println("readyok")
	case "ucinewgame":
		p.Prepare()
		p.newGame()
	case "position":
		if p.brain == nil {
			p.Prepare()
			p.newGame()
		}
		p.parsePosition(command)
	case "debug":
	case "register":
	case "go":
		p.brain.Search(engine.Condition{})
	case "ponderhit":
		//Rival goes the same step of the engine thought
		//so stop thinking and give the answer
	case "stop":
		p.brain.Stop()
	//From now on are user-defined commands
	case "board":
		p.brain.ShowBoard()
	}
	return nil
}

func (p *ProtocolUCI) setOption(name string, value string) (err error) {
	if v, ok := engine.Options[name]; ok {
		switch v.(type) {
		case int:
			engine.Options[name], _ = strconv.Atoi(value)
		case float64:
			engine.Options[name], _ = strconv.ParseFloat(value, 64)
		case bool:
			engine.Options[name], _ = strconv.ParseBool(value)
		case string:
			engine.Options[name] = value
		}
		err = engine.CheckOptions()
	} else {
		err = fmt.Errorf("option \"%s\" not exist", name)
	}
	return
}

func (p *ProtocolUCI) info(msg ...string) {
	println("info string", msg)
}

func (p *ProtocolUCI) newGame() {
	p.brain.NewGame()
}

func (p *ProtocolUCI) parsePosition(command string) {
	// position [fen <fenstring> | startpos ]  moves <move1> .... <movei>
	command = strings.TrimSpace(strings.TrimPrefix(command, "position"))
	parts := strings.Split(command, "moves")
	if len(command) == 0 || len(parts) > 2 {
		err := fmt.Errorf("%v wrong length=%v", parts, len(parts))
		p.info(err.Error())
		return
	}
	spliced := strings.Split(parts[0], " ")
	spliced[0] = strings.TrimSpace(spliced[0])
	switch spliced[0] {
	case "startpos":
		parts[0] = chess.StartingPosition().String()
	case "fen":
		parts[0] = strings.TrimSpace(strings.TrimPrefix(parts[0], "fen"))
	default:
		err := fmt.Errorf("%#v must be %#v or %#v", spliced[0], "fen", "startpos")
		p.info(err.Error())
		return
	}
	err := p.brain.SetPosition(parts[0])
	if err != nil {
		p.info(parts[0])
		return
	}
	if len(parts) == 2 {
		parts[1] = strings.ToLower(strings.TrimSpace(parts[1]))
		moves := strings.Fields(parts[1])
		for _, move := range moves {
			err := p.brain.Move(move)
			if err != nil {
				p.info(move, " is not a correct move")
				return
			}
		}
	}
}
