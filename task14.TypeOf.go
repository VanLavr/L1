package main

import (
	"fmt"
	"reflect"
)

func main() {
	var (
		a chan int
		b string
		c int
		d bool
	)

	inspectType(a)
	inspectType(b)
	inspectType(c)
	inspectType(d)

	inspectType2(a)
	inspectType2(b)
	inspectType2(c)
	inspectType2(d)
}

func inspectType(i any) {
	fmt.Println(reflect.TypeOf(i))
}

func inspectType2(i any) {
	switch i.(type) {
	case chan int:
		fmt.Println("chan int")
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	}
}
