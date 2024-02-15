package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	fmt.Print("Anna sana: ")
	input := ReadInput()
	inputLength := len(input)
	if (inputLength > 1) {
		fmt.Printf("Sanassa %s on %d kirjainta\n", input, inputLength)
	}
	fmt.Println("Kiitos!")
}
