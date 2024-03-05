package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

func main() {
	// getting the number of workers from provided argument
	sWorkersNumber := os.Args[1]
	// converting it to int
	workersNumber, err := strconv.Atoi(sWorkersNumber)
	if err != nil || workersNumber < 1 {
		log.Fatal(err)
	}

	// creating context with cancelation to control the workers and publisher
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	// when cancel function is called all the goroutines recieves the Done channel closure => than handling it
	// and releasing resources
	defer cancel()

	// waitgroup stands for making the main goroutine wait for all another goroutines to stop
	var wg sync.WaitGroup
	// this is the main thread stands for interacting between publisher and workers
	mainChannel := make(chan int)

	// spawning the workers...
	for i := 0; i < workersNumber; i++ {
		// incrementing wg delta for all of the workers
		wg.Add(1)
		go func(i int) {
			Worker(ctx, mainChannel, i, &wg)
		}(i)
	}

	// spawning the publisher
	// incrementing wg delta for publisher
	wg.Add(1)
	go Publisher(ctx, mainChannel, &wg)

	// waiting for all spawned goroutines to stop
	wg.Wait()
	// waiting for cancelling the context
	<-ctx.Done()
	fmt.Println("\nmain has stopped")
}

func Publisher(ctx context.Context, main chan<- int, wg *sync.WaitGroup) {
	// decrementing wg delta
	defer wg.Done()
	for {
		select {
		// generating messages to the channel
		case main <- rand.Int():
			time.Sleep(time.Millisecond * 350)
		// cancelling the goroutine
		case <-ctx.Done():
			fmt.Println("\npublisher has stopped")
			return
		}
	}
}

func Worker(ctx context.Context, main <-chan int, number int, wg *sync.WaitGroup) {
	// decrementing wg delta
	defer wg.Done()
	for {
		select {
		// reading message from main channel and printing it to stdout
		case val, ok := <-main:
			if !ok {
				log.Fatal("can not read from channel")
				return
			}
			fmt.Printf("worker %d read: %d\n", number, val)
		// cancelling the goroutine
		case <-ctx.Done():
			fmt.Printf("\nworker %d has stopped\n", number)
			return
		}
	}
}
