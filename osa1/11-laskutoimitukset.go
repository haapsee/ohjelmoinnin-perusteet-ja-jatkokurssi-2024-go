package main

import (
	"fmt"
)

func main() {
	// 27 + 15 = 42
	// 27 - 15 = 12
	// 27 * 15 = 405
	// 27 / 15 = 1.8
	x := 27
	y := 15
	fmt.Printf("%d + %d = %d\n", x, y, (x+y))
	fmt.Printf("%d - %d = %d\n", x, y, (x-y))
	fmt.Printf("%d * %d = %d\n", x, y, (x*y))
	fmt.Printf("%d / %d = %g\n", x, y, (float64(x)/float64(y)))
}
