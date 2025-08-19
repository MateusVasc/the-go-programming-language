package main

import (
	"fmt"

	osargs "github.com/MateusVasc/the-go-programming-language/os-args"
)

func main() {
	fmt.Println("The Go Programming Language")
	fmt.Println("")

	// 1.2
	fmt.Println("1.2 Command-Line Arguments")
	osargs.Echo()
	osargs.RangeEcho()
	osargs.InvokedEcho()
	osargs.EachEcho()
}
