package main

import (
	"fmt"
	"math/rand"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// justString = v[:100]
	justString = string(CopyRunes([]rune(v), 100))
}

func main() {
	someFunc()
	fmt.Println(justString)
}

func createHugeString(n int) string {
	// const letters = "qwertyuiopasdfghjklzxcvbnm"
	const letters = "абвгдеёжзийклмнопрстуфхцчшщъыьэюя"
	runes := []rune(letters)
	buffer := make([]rune, n)
	for b := range buffer {
		buffer[b] = runes[rand.Intn(len(runes))]
	}

	return string(buffer)
}

func CopyRunes(source []rune, n int) []rune {
	result := make([]rune, n)
	for i := 0; i < n; i++ {
		result[i] = source[i]
	}
	return result
}
