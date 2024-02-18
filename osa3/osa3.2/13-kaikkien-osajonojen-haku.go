package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Sana: ")
	word, _ := r.ReadString('\n')
	word = strings.TrimSpace(word)

	fmt.Print("Merkki: ")
	char, _ := r.ReadString('\n')
	char = strings.TrimSpace(char)

	for i := 0; i < len(word); i++ {
		if string(word[i]) == char {
			fmt.Println(word[i:min(i+3, len(word))])
		}
	}
}
