package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("Luku: ")
	input, _ := r.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if num % 3 == 0 {
		fmt.Print("Fizz")
	}
	if num % 5 == 0 {
		fmt.Print("Buzz")
	}
	if num % 3 == 0 || num % 5 == 0 {
		fmt.Println("")
	}
}
