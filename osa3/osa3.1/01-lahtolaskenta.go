package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadNumber(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	value, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	return int(value)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna luku: ")
	for i := ReadNumber(reader); i > 0; i -= 1 {
		fmt.Printf("%d\n", i)
	}
	fmt.Println("Nyt!")
}
