package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Sana: ")
	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)

	for i := 0; i < 20 - len(input); i++ {
		fmt.Print("*")
	}
	fmt.Println(input)
}
