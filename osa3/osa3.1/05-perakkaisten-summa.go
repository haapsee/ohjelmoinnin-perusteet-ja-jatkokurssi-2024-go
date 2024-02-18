package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Mihin asti: ")
	input, _ := reader.ReadString('\n')
	n, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	s := 0

	for i := 0; s < int(n)+1; i++ {
		s += i
	}
	fmt.Printf("%d\n", s)
}
