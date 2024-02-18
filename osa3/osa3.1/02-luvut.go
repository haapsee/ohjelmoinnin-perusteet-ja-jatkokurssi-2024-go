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
	value, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	for i := 1; i < int(value); i++ {
		fmt.Printf("%d\n", i)
	}
}
