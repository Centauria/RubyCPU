package engine

import (
	"fmt"
	"sync"
	"time"
)

type Engine struct {
	ThreadsNum int
	threads    []thread
	Stdout     chan string
	mu         *sync.Mutex
	wg         *sync.WaitGroup
}

type thread struct {
	Engine *Engine
	Run    func()
	Stdin  chan string
	wg     *sync.WaitGroup
}

func NewEngine(threadsNum int) *Engine {
	wg := &sync.WaitGroup{}
	engine := &Engine{
		ThreadsNum: threadsNum,
		threads:    make([]thread, threadsNum),
		Stdout:     make(chan string, threadsNum),
		mu:         &sync.Mutex{},
		wg:         wg,
	}
	for i := 0; i < threadsNum; i++ {
		wg.Add(1)
		t := thread{
			Engine: engine,
			Stdin:  make(chan string),
			wg:     wg,
		}
		t.Run = func() {
			defer wg.Done()
			cmd := <-t.Stdin
			time.Sleep(1 * time.Second)
			engine.Stdout <- fmt.Sprintf("Reply cmd(%s)", cmd)
		}
		engine.threads[i] = t
	}
	return engine
}

func (e Engine) Start() {
	for _, t := range e.threads {
		go t.Run()
	}
	e.wg.Wait()
	for range e.threads {
		println(<-e.Stdout)
	}
}

func (e Engine) Input(cmd string) {
	for _, t := range e.threads {
		t.Input(cmd)
	}
}

func (t thread) Input(cmd string) {
	t.Stdin <- cmd
}
