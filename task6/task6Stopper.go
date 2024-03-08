package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// 1 channel
// 2 context
// 3 atomic
// 4 mutex (very similar to atomic)

type Stopper struct {
	sync.RWMutex
	isStopped bool
}

func (s *Stopper) IsStopped() bool {
	s.RLock()
	defer s.RUnlock()
	return s.isStopped
}

func (s *Stopper) Stop() {
	s.Lock()
	defer s.Unlock()
	s.isStopped = true
}

func main() {
	// 1:
	done := make(chan struct{})

	go func(done <-chan struct{}) {
		for {
			select {
			default:
				time.Sleep(time.Millisecond * 300)
				fmt.Println("goroutine is running")
			case <-done:
				fmt.Println("goroutine stopped")
				return
			}
		}
	}(done)

	time.Sleep(time.Second * 2)
	done <- struct{}{}

	// 2:
	ctx, close := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			default:
				time.Sleep(time.Millisecond * 300)
				fmt.Println("goroutine is running")
			case <-ctx.Done():
				fmt.Println("goroutine stopped")
				return
			}
		}
	}(ctx)

	time.Sleep(time.Second * 2)
	close()

	// 3 and 4 are bad, because (golang official dosc): Share memory by communicating; don't communicate by sharing memory
	// 3:
	var commutator int32 = 1

	go func(commutator *int32) {
		for *commutator == 1 {
			fmt.Println("goroutine is running")
			time.Sleep(time.Millisecond * 300)
		}
		fmt.Println("goroutine stopped")
	}(&commutator)

	time.Sleep(time.Second * 2)
	atomic.AddInt32(&commutator, 1) // or: atomic.StoreInt32(&commutator, 2)

	// 4:
	var stopper Stopper

	go func(stop *Stopper) {
		for !stop.IsStopped() {
			fmt.Println("goroutine is running")
			time.Sleep(time.Millisecond * 300)
		}
		fmt.Println("goroutine stopped")
	}(&stopper)

	time.Sleep(time.Second * 2)
	stopper.Stop()
	time.Sleep(time.Second)
}
