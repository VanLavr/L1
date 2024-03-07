package main

import (
	"fmt"
)

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(arr)
	set := make(map[string]struct{})
	for _, v := range arr {
		set[v] = struct{}{}
	}
	PrintSet(set)
}

func PrintSet(m map[string]struct{}) {
	for k := range m {
		fmt.Printf("%s ", k)
	}
	fmt.Println()
}
