package worker

import (
	"context"
	"sync"
)

type WorkerFunc func(ctx context.Context) error

type Group struct {
	wg sync.WaitGroup
}

func (g *Group) Go(ctx context.Context, fn WorkerFunc) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		_ = fn(ctx)
	}()
}

func (g *Group) Wait() {
	g.wg.Wait()
}
