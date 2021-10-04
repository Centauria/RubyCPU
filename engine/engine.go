package engine

import (
	"fmt"
	"sync"
	"time"
)

type Engine struct {
	daemon  *thread
	channel chan string
	mu      *sync.Mutex
}

type thread struct {
	Engine *Engine
	Run    func()
}

func NewEngine() *Engine {
	engine := &Engine{
		channel: make(chan string),
		mu:      &sync.Mutex{},
	}
	t := &thread{
		Engine: engine,
	}
	t.Run = func() {
		cmd := <-engine.channel
		time.Sleep(1 * time.Second)
		engine.channel <- fmt.Sprintf("Reply cmd(%s)", cmd)
	}
	engine.daemon = t
	return engine
}

func (e *Engine) Start() {
	defer close(e.channel)
	go e.daemon.Run()
	println(<-e.channel)
}

func (e *Engine) Input(cmd string) {
	e.channel <- cmd
}
