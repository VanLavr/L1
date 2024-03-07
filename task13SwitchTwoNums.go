package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(a, b)

	b -= a
	a += b
	b = a - b

	// or: a, b = b, a

	fmt.Println(a, b)
}
