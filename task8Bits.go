package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"strconv"
)

func main() {
	// example
	var a int = 117
	fmt.Printf("%b\n", a)
	a = a &^ (1 << 5)
	fmt.Printf("%b\n\n", a)

	// new scanner (fmt.Scan() do not flush input from \n after reading)
	s := bufio.NewReader(os.Stdin)

	var n int64

	fmt.Println("enter a number:")
	stringn, err := s.ReadString('\n')
	// flushing input from \n by hand (cast string to []rune, than strip slice, removing two last elements, than cast to string)
	stringn = string([]rune(stringn[:len(stringn)-1]))
	if err != nil {
		log.Fatal(err)
	}
	n, err = strconv.ParseInt(stringn, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("your num is: %b (%d)\n", n, n)

	var bitPos int
	fmt.Println("Which bit do you want to change? (count from left to right, from 1 to ...)")
	stringBitPos, err := s.ReadString('\n')
	// flushing input from \n by hand (cast string to []rune, than strip slice, removing two last elements, than cast to string)
	stringBitPos = string([]rune(stringBitPos[:len(stringBitPos)-1]))
	if err != nil {
		log.Fatal(err)
	}
	bitPos, err = strconv.Atoi(stringBitPos)
	if err != nil {
		log.Fatal(err)
	}

	var isZero int
	fmt.Println("Set it to 0 or to 1?")
	stringIsZero, err := s.ReadString('\n')
	// flushing input from \n by hand (cast string to []rune, than strip slice, removing two last elements, than cast to string)
	stringIsZero = string([]rune(stringIsZero[:len(stringIsZero)-1]))
	if err != nil {
		log.Fatal(err)
	}
	isZero, err = strconv.Atoi(stringIsZero)
	if err != nil {
		log.Fatal(err)
	}

	if isZero == 0 {
		n = n &^ (1 << (bitPos - 1))
	} else {
		n = n | (1 << (bitPos - 1))
	}
	fmt.Printf("now your number is: %b (%d)\n", n, n)
}

// how is this working:
/*



HOW TO SET i'th BIT TO 1:

our number is: 44 (for example) and our bit is second bit

1) we converting it to binary format (in out head, not in programm):
   44 => 101100

2) now we are creating a number with only one 1 bit in exactly that place where we want
   to see the 1 in out number, this way:
   1 << bitPosition ("bitPosition" is that place, and "<<" is binary shift to left
   (multiply by two if bitPosition = 1, and multiply by four if bitPosition = 2 and so on... in other words))
   so now we have two numbers:
   101100 and 100 (we have to set bitPosition to 4 to move the 1 like this: 1 -> 10 -> 100 -> 1000 -> 10000)

3) now we can simply OR this numbers and get third bit as one:
   101100 |
   010000
   ------
   111100

   finally: we done it


HOW TO SET i'th BIT TO 0:

our number is: 44 (for example) and our bit is fifth bit

1) we converting it to binary format (in out head, not in programm):
   44 => 101100

2) now we are creating a number with only one 1 bit in exactly that place where we want
   to see the 1 in out number, this way:
   1 << bitPosition ("bitPosition" is that place, and "<<" is binary shift to left
   (multiply by two if bitPosition = 1, and multiply by four if bitPosition = 2 and so on... in other words))
   so now we have two numbers:
   101100 and 100 (we have to set bitPosition to 2 to move the 1 like this: 1 -> 10 -> 100)

3) now we can simply NAND (not and) this numbers and get third bit as one:
   101100 &^
   000100
   ------
   101000

   finally: we done it

*/
