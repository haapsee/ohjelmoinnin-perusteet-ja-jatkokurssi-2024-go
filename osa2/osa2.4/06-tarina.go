package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func ReadString(r *bufio.Reader) string {
	fmt.Print("Anna sana: ")
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	story := ""
	lastWord := ReadString(reader)

	for lastWord != "loppu" {
		story = fmt.Sprintf("%s %s", story, lastWord)
		nextWord := ReadString(reader)
		if nextWord == lastWord {
			break
		}
		lastWord = nextWord
	}

	fmt.Println(story[1:])
}
