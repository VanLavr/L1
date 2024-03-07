package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}

	index := sort.Search(len(arr), func(i int) bool { return arr[i] >= 13 })

	if index < len(arr) && arr[index] == 13 {
		fmt.Println("Found 13 at index:", index)
	} else {
		fmt.Println("13 not found in the slice")
	}
}
