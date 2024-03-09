package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// 1:
	fmt.Println("aaa" + "bbb") // creates a new string every time its called

	// 2:
	var b strings.Builder // do not create a new string // It minimizes memory copying. (docs)
	b.WriteString("sss")
	b.WriteString("ddd")
	fmt.Println(b.String())

	// 3:
	s := []string{"hello", ",", "world"} // create a new string
	fmt.Println(strings.Join(s, " "))

	// 4:
	var buffer bytes.Buffer
	buffer.WriteString("Hello, ")
	buffer.WriteString("Gophers!")
	fmt.Println(buffer.String())

	// 5:
	a := "ads"
	c := "da"
	fmt.Sprintf("%s%s", a, c)
}

func Plus(a, b string) string {
	return a + b
}

func Builder(a, b string) string {
	var bl strings.Builder
	bl.WriteString(a)
	bl.WriteString(b)
	return bl.String()
}

func Join(arr []string, sep string) string {
	return strings.Join(arr, sep)
}

func Buffer(a, b string) string {
	var bf bytes.Buffer
	bf.WriteString(a)
	bf.WriteString(b)
	return bf.String()
}

func Sprint(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

/*

   Use the strings.Builder type instead of concatenating strings with the + operator or fmt.Sprintf()

   Pre-allocate the necessary space for the strings.Builder to avoid unnecessary allocations

   If concatenating many strings at once, consider using the strings.Join() function instead of a loop with strings.Builder.

   Be mindful of the number of strings being concatenated, as excessive concatenation can cause performance issues.


goos: linux
goarch: amd64
pkg: github.com/VanLavr/L1/task1.1
cpu: AMD Ryzen 7 4800H with Radeon Graphics
BenchmarkPlus-16        55658852               22.44 ns/op
BenchmarkBuilder-16      7568839               161.6 ns/op
BenchmarkJoin-16         9731810               125.3 ns/op
BenchmarkBuffer-16       6570566               186.0 ns/op
BenchmarkSprint-16       2814986               412.0 ns/op
PASS
ok      github.com/VanLavr/L1/task1.1   7.660s
*/
