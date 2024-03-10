package main

func main() {
	// slices
	var s []string   // nil slice
	s1 := []string{} // var s1 = []string{}, slice literal
	s2 := make([]string, 0, 0)
	s3 := *new([]string)
	arr := [...]string{"s", "s", "s"}
	s4 := arr[:]

	// maps
	var m map[int]int   // nil map
	m1 := map[int]int{} // map literal
	m2 := make(map[int]int, 0)
	m3 := *new(map[int]int)
}
