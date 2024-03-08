package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Iterator struct {
	sync.Mutex
	i int
}

func (i *Iterator) Add() {
	i.Lock()
	defer i.Unlock()
	i.i++
}

func (i *Iterator) Get() int {
	i.Lock()
	defer i.Unlock()

	return i.i
}

func main() {
	max := rand.Intn(100)
	fmt.Println("expected result:", max)

	I := new(Iterator)

	var wg sync.WaitGroup
	for i := 0; i < max; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			I.Add()
		}()
	}

	wg.Wait()

	fmt.Println("result:", I.Get())
}
