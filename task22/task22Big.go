package main

import (
	"fmt"
	"math/big"
)

func main() {
	a, _ := new(big.Int).SetString("1000000000000000000000000000000000000000000000000", 10)
	b, _ := new(big.Int).SetString("500000000000000000000000000000000000000000000000", 10)
	c := new(big.Int)

	fmt.Println("a =", a, ";", " b =", b, ";", " c =", c)

	fmt.Println("a * b =", c.Mul(a, b))
	fmt.Println("a / b =", c.Div(a, b))
	fmt.Println("a + b =", c.Add(a, b))
	fmt.Println("a - b =", c.Sub(a, b))
}
