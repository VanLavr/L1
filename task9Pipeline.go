package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 5, 6}

	arrChann := ArrToChann(arr)
	squares := Squarer(arrChann)
	for square := range squares {
		fmt.Println(square)
	}
}

func ArrToChann(arr []int) <-chan int {
	arrch := make(chan int)
	go func() {
		defer close(arrch)
		for _, num := range arr {
			arrch <- num
		}
	}()
	return arrch
}

func Squarer(arrch <-chan int) <-chan int {
	sqrs := make(chan int)
	go func() {
		defer close(sqrs)
		for num := range arrch {
			sqrs <- num * num
		}
	}()
	return sqrs
}
