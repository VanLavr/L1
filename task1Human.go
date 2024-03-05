package main

import "fmt"

// Parent structure
type Human struct {
	Name string
	Age  int
}

// Method 1
func (h Human) Talk(speech string) {
	fmt.Println(speech)
}

// Method 2
func (h Human) SelebrateBirthday() {
	fmt.Printf("%s have selebrated %d'th birthday!\n", h.Name, h.Age)
}

// Method 3
func (h Human) Walk(distance int) {
	fmt.Printf("Walking for %d meters\n", distance)
}

// Structure "inharitor" (anonimous embedding or composition)
type Action struct {
	Human
}

func main() {
	h := new(Human)
	h.Name = "Joch"
	h.Age = 35

	// Prints text with name and age
	h.SelebrateBirthday()
	// Prints text with number of meters
	h.Walk(123)

	act := new(Action)
	act.Name = "Yanni"
	act.Age = 34

	// Prints text
	act.Talk("Hello, world!")
	// Prints text with name and age
	act.SelebrateBirthday()
	// Prints the text with number of meters
	act.Walk(123)
}
