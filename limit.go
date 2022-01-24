package glimit

import "sync"

type Limit struct {
	c  chan struct{}
	wg sync.WaitGroup
}

func New(n int) *Limit {
	return &Limit{
		c:  make(chan struct{}, n),
		wg: sync.WaitGroup{},
	}
}

func (g *Limit) Add() {
	g.c <- struct{}{}
	g.wg.Add(1)
}

func (g *Limit) Done() {
	<-g.c
	g.wg.Done()
}

func (g *Limit) Wait() {
	g.wg.Wait()
}
