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
	spaces := 28 - len(input)
	spaces = spaces / 2

	for i := 0; i < 29; i++ {
		fmt.Print("*")
	}
	fmt.Println("*")

	fmt.Print("*")
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	fmt.Print(input)
	for i := 0; i < spaces + len(input)%2; i++ {
		fmt.Print(" ")
	}
	fmt.Println("*")

	for i := 0; i < 30; i++ {
		fmt.Print("*")
	}
	fmt.Println("")
}
