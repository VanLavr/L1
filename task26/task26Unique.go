package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "helLo word!"
	fmt.Println(IsUnique(s))
}

func IsUnique(s string) bool {
	sdata := strings.ToLower(s)

	m := make(map[rune]struct{})
	data := []rune(sdata)

	for _, r := range data {
		_, found := m[r]
		if !found {
			m[r] = struct{}{}
		} else {
			return false
		}
	}

	return true
}
