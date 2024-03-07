package main

import (
	"fmt"
	"sort"
)

// create an alias for []int to make it implement sort.Interface
type sortInt []int

func (a sortInt) Len() int           { return len(a) }
func (a sortInt) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortInt) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	arr := sortInt{123, 3, 24, -123, -3, 4, 32, -4, 39, 4013, 1, 94, -987, 890, 84}
	sort.Sort(arr) // sorts a slice via quicksort
	fmt.Println(arr)
}
