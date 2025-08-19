package osargs

import (
	"fmt"
	"os"
	"strings"
)

func Echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func RangeEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func JoinEcho() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

// Exercise 1.1: Modify the echo program to also print os.Args[0], the name of the command that invoked it.
func InvokedEcho() {
	fmt.Println(strings.Join(os.Args[:], " "))
}

// Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
func EachEcho() {
	for i, arg := range os.Args[:] {
		fmt.Printf("Index: %d - Value: %s\n", i, arg)
	}
}

// Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient versions and the one that
// uses strings.Join. (Section1.6 illustrates part of the time package, and Section 11.4 shows how to write benchmark tests for
// systematic performance evaluation.)
