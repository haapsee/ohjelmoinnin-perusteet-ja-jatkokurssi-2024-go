package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Anna merkkijono: ")
	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)

	if (len(input) < 1) {return;}
	fmt.Printf("\n%s\n", input)
	for i := 0; i < len(input); i++ {
		fmt.Print("-")
	}
	fmt.Println("")
}
