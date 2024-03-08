package main

import (
	"fmt"
	"slices"
)

func main() {
	// first solution:
	a := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(a, "\nenter an index of element to delete:")
	var i int
	fmt.Scanf("%d", &i)
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a, "\n")

	// second solution:
	b := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(b, "\nenter an index of element to delete:")
	var k int
	fmt.Scanf("%d", &k)
	slices.Delete(b, k, k+1) // it will leave the slice with the same lenght so 0 at the end appears
	fmt.Println(b, "\n")

	// third solution:
	c := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(c, "\nenter an element to delete:")
	var j int
	fmt.Scanf("%d", &j)
	slices.DeleteFunc(c, func(elem int) bool { // it will also leave the slice with the same lenght so 0 at the end appears
		if elem == j {
			return true
		}
		return false
	})
	fmt.Println(c)
}
