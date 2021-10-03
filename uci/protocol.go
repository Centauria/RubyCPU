package uci

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
