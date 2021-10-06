package engine

type Flag int8

const (
	Exact Flag = iota
	Low
	High
)

type transTableEntry struct {
	key      uint64
	depth    int8
	typeFlag Flag
	score    float32
}

//`HashSize` defined as length of hash table temporarily
