package main

import "fmt"

func Intersect(set1, set2 map[string]struct{}) map[string]struct{} {
	result := make(map[string]struct{})

	for k := range set1 {
		for j := range set2 {
			if k == j {
				result[k] = struct{}{}
			}
		}
	}

	return result
}

func SetPrinter(m map[string]struct{}) {
	for k := range m {
		fmt.Printf("%s ", k)
	}

	fmt.Println()
}

func main() {
	set1 := make(map[string]struct{})
	set2 := make(map[string]struct{})

	set1["adf"] = struct{}{}
	set1["asdf"] = struct{}{}
	set1[";lkj"] = struct{}{}
	set1["zxcv"] = struct{}{}

	set2["zxcv"] = struct{}{}
	set2["asdf"] = struct{}{}
	set2["1"] = struct{}{}
	set2["2"] = struct{}{}

	intersection := Intersect(set1, set2)
	SetPrinter(intersection)
}
