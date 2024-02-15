package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func MinMax(x int64, y int64) int64 {
	if (x > y) {
		return x
	}
	return y
}

func main() {
	fmt.Print("Anna ensimmäinen luku: ")
	x, _ := strconv.ParseInt(ReadInput(), 10, 64)
	fmt.Print("Anna toinen luku: ")
	y, _ := strconv.ParseInt(ReadInput(), 10, 64)
	if (x == y) {
		fmt.Println("Luvut ovat yhtä suuret!")
	} else {
		fmt.Printf("Suurempi luku: %d\n", MinMax(x, y))
	}
}
