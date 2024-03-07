package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		a chan int
		b string
		c int64
		d uint64
		e uintptr
		f any
		g []rune
	)

	inspectType(a)
	inspectType(b)
	inspectType(c)
	inspectType(d)
	inspectType(e)
	inspectType(f)
	inspectType(g)
}

func inspectType(i any) {
	fmt.Println(reflect.TypeOf(i))
}
