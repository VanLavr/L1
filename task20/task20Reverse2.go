package main

import (
	"fmt"
	"strings"
)

func main() {
	sample := "snow dog sun"

	fmt.Printf("[%s] ---> [%s]\n", sample, ReverseWords(sample))
}

func ReverseWords(s string) string {
	words := strings.Split(s, " ")

	var result string

	for i, _ := range words {
		result += words[len(words)-i-1]
		result += " "
	}
	result = result[:len(result)-1]

	return result
}
