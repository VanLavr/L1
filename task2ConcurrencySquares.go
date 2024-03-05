package main

import (
	"fmt"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	// first solution (unordered)
	arr := []int{2, 4, 6, 8, 10}
	for _, val := range arr {
		go Square(val)
	}
	time.Sleep(time.Millisecond * 300)

	// second solution (unordered)
	fmt.Println()
	wg := sync.WaitGroup{}

	for _, val := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Square(val)
		}()
	}

	wg.Wait()

	// third solution (ordered)
	fmt.Println()

	for _, val := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			Square(val)
		}()
		wg.Wait()
	}

	// fourth solution (ordered)
	fmt.Println()
	data := make(chan int)
	defer close(data)
	done := make(chan struct{})
	defer close(done)
	go ChanSquare(data, done)
	for _, val := range arr {
		data <- val
	}
	done <- struct{}{}                 // or close(done)
	time.Sleep(time.Millisecond * 300) // or use waitgroup

	// fifth solution (ordered)
	fmt.Println()
	wg.Add(1)
	go ChanSquareWg(data, done, &wg)
	for _, val := range arr {
		data <- val
	}
	close(done)
	wg.Wait()

	// http.ListenAndServe("localhost:8989", nil)
}

func Square(arg int) {
	defer fmt.Println("goroutine stopped")
	fmt.Println(arg * arg)
}

func ChanSquare(data <-chan int, done <-chan struct{}) {
	defer fmt.Println("goroutine stopped")
	for {
		select {
		case val, ok := <-data:
			if !ok {
				fmt.Errorf("cannot read from channel")
			}
			fmt.Println(val * val)
		case <-done:
			return
		}
	}
}

func ChanSquareWg(data <-chan int, done <-chan struct{}, wg *sync.WaitGroup) {
	defer func() {
		fmt.Println("goroutine stopped")
		wg.Done()
	}()
	for {
		select {
		case val, ok := <-data:
			if !ok {
				fmt.Errorf("cannot read from channel")
			}
			fmt.Println(val * val)
		case <-done:
			return
		}
	}
}
