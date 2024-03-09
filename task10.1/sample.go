package main

import (
	"fmt"
	"unsafe"
)

// changes local pointer
func update(p *int) {
	b := 2
	p = &b
}

func unsafeUpdate(p unsafe.Pointer) {
	*((*int)(p)) = 42
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p)
	unsafeUpdate(unsafe.Pointer(p))
	fmt.Println(*p)
}

// argument p is overwritten before first use (SA4009)
// sample.go:7:2: assignment to p go-staticcheck
