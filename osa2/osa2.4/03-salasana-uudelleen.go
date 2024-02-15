package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func NewStringReader() func() string {
	reader := bufio.NewReader(os.Stdin)

	return func() string {
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}
}

func main() {
	reader := NewStringReader()

	fmt.Print("Salasana: ")
	password := reader()
	fmt.Print("Toista salasana: ")

	for password != reader() {
		fmt.Println("Ei ollut sama!")
		fmt.Print("Toista salasana: ")
	}
	fmt.Println("Käyttäjätunnus luotu!")
}
