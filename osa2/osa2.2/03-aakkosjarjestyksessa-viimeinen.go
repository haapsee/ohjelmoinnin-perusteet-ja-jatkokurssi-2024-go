package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func Compare(a string, b string) int {
	length := len(a)
	if (length > len(b)) {
		length = len(b)
	}
	if (a == b) {
		return 0
	}
	for i := 0; i < length; i++ {
		if (a[i] < b[i]) {
			return 1
		} else if (b[i] < a[i]) {
			return 2
		}
	}
	if (len(a) < len(b)) {
		return 1
	}
	return 2
}

func ReadInput(reader *bufio.Reader) string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna 1. sana: ")
	w1 := ReadInput(reader)
	fmt.Print("Anna 2. sana: ")
	w2 := ReadInput(reader)

	result := Compare(w1, w2)

	if (result == 0) {
		fmt.Println("Annoit saman sanan kahdesti.")
	} else if (result == 1) {
		fmt.Printf("%s on aakkosjärjestyksessä viimeinen.\n", w2)
	} else if (result == 2) {
		fmt.Printf("%s on aakkosjärjestyksessä viimeinen.\n", w1)
	}
}
