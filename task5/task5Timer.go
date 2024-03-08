package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// reading N seconds from cli args
	n := os.Args[1]
	N, err := strconv.Atoi(n)
	if err != nil || N < 1 {
		log.Fatal(err)
	}

	// creating channel for data and shutting down
	data := make(chan int)
	stop := make(chan struct{})

	// running a timer
	go func() {
		time.Sleep(time.Second * time.Duration(N))
		close(stop)
	}()

	// running a goroutine that will write data to the channel
	go func() {
		for {
			select {
			case data <- rand.Int():
				time.Sleep(time.Millisecond * 300)
			case <-stop:
				close(data)
				return
			}
		}
	}()

	// reading data from the channel
	for val := range data {
		fmt.Println(val)
	}
}
