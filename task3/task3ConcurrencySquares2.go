package main

import (
	"fmt"
	"sync"
)

// square => summ => print
func main() {
	arr := []int{2, 4, 6, 8, 10}

	// first solution with "pipeline"
	var s int

	// stage one - make channel from array
	dataChannel := arrToChannel(arr)
	// stage two - square the numbers
	squares := squarer(dataChannel)
	// stage three - make a summ
	for num := range squares {
		s += num
	}
	fmt.Println(s)

	// second solution
	fmt.Println()

	wg := sync.WaitGroup{}

	data := make(chan int)
	defer close(data)
	done := make(chan struct{})

	// asynchronous summator waits for the data in the channel and increments the result
	// main goroutine have to wait (with wg.Wait()) till this function prints the result
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(summator(data, done))
	}()

	// asynchronous senders (to the channel that summator recieves) are created here
	// they are squaring the number from the array asynchronously and sending it to the channel
	wg1 := sync.WaitGroup{}
	for _, num := range arr {
		wg1.Add(1)
		go func() {
			defer wg1.Done()
			data <- squareNum(num)
		}()
	}
	wg1.Wait()

	close(done)
	wg.Wait()
}

func squareNum(num int) int {
	return num * num
}

func summator(data <-chan int, done <-chan struct{}) int {
	var sum int
	for {
		select {
		case val, ok := <-data:
			if !ok {
				fmt.Println("can not read from channel")
				return 0
			}
			sum += val
		case <-done:
			return sum
		}
	}
}

func squarer(data <-chan int) <-chan int {
	squares := make(chan int)
	go func() {
		defer close(squares)
		for num := range data {
			squares <- num * num
		}
	}()
	return squares
}

func arrToChannel(arr []int) <-chan int {
	data := make(chan int)
	go func() {
		defer close(data)
		for _, num := range arr {
			data <- num
		}
	}()
	return data
}
