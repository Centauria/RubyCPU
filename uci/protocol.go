package uci

import "context"

type Protocol interface {
	Prepare()
	Process()
	Cleanup()
}

type ProtocolUCI struct {
}

func (p ProtocolUCI) Prepare() {

}

func (p ProtocolUCI) Process() {

}

func (p ProtocolUCI) Cleanup() {

}

func (p ProtocolUCI) Handle(ctx context.Context, command string) error {
	switch command {
	case "uci":
		println("uciok")
	}
	return nil
}
