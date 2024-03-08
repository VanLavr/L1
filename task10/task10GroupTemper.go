package main

import (
	"fmt"
	"log"
)

func main() {
	temps := []float32{-30.0, 0.0, 4.2, -39.9, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 20.0}
	groups := make(map[int][]float32)

	// finding min and max
	min := temps[0]
	max := temps[0]
	for _, tmp := range temps {
		if tmp > max {
			max = tmp
		}
		if tmp < min {
			min = tmp
		}
	}

	// defining start for iterator which will help to create groups (-39.9 => -30 and 32.5 => 30)
	start := int(min) / 10 * 10
	finish := int(max) / 10 * 10

	// creating groups
	for i := start; i < finish+1; i += 10 {
		groups[i] = make([]float32, 0)
	}

	for k := range groups {
		for _, tmp := range temps {
			// if temperature above the zero or equal than x <= TEMP < x + 10
			if tmp >= 0.0 {
				if int(tmp) < k+10 && int(tmp) > k-1 {
					group, found := groups[k]
					if !found {
						log.Fatal("something went wrong")
					}
					// adding temperature to the group
					group = append(group, tmp)
					groups[k] = group
				}

				// if temperature below the zero than x - 10 < TEMP <= x
			} else {
				if int(tmp) < k+1 && int(tmp) > k-10 {
					group, found := groups[k]
					if !found {
						log.Fatal("something went wrong")
					}
					// adding temperature to the group
					group = append(group, tmp)
					groups[k] = group
				}
			}
		}
	}

	// deleting all groups that have no data within
	for k := range groups {
		group, found := groups[k]
		if !found {
			log.Fatal("something went wrong")
		}
		if len(group) == 0 {
			delete(groups, k)
		}
	}

	for k, v := range groups {
		fmt.Println(k, v)
	}
}
