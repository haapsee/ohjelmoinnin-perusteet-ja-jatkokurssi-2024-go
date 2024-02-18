package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna merkkijono: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	for i := len(input); i > 0; i-- {
		fmt.Printf("%s\n", string(input[i-1]))
	}
}
