package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func ReadInput(r *bufio.Reader) bool {
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input) != "ei"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	running := true

	for running {
		fmt.Print("moi\nJatketaanko? ")
		running = ReadInput(reader)
	}

	fmt.Println("ei sitten")
}
