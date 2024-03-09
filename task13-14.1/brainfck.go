package main

import "fmt"

func someAction(v []int8, b int8) {
	v[0] = 100
	v = append(v, b)
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a) // last element has changed because slices if you copy a slice in golang and change the "child" it will affect to "parent" and in reverse

	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)

	fmt.Print(slice) // last element has not changed because "append()" function returns a NEW slice and we reassigning this new slice to previous variable,
	// so now slice in anonimous funcion is not a "child" for slice in main function anymore
}
