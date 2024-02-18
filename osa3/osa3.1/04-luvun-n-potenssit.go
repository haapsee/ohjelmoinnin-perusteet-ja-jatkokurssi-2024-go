package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadNum(reader *bufio.Reader) int {
	input, _ := reader.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	return int(num)
}

func Pows(reader *bufio.Reader) func() {
	fmt.Print("Mihin asti: ")
	x := ReadNum(reader)+1

	return func() {
		fmt.Print("Mikä kerroin: ")
		n := ReadNum(reader)

		for i := 1; i < x; i *= n {
			fmt.Printf("%d\n", i)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	pows := Pows(reader)
	pows()
}
