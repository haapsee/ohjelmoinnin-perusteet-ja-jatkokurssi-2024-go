package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("anna merkkijono: ")
	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)

	for i := 0; i < len(input); i++ {
		fmt.Println(input[0:i+1])
	}
}
