package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func main() {
	fmt.Print("Anna osa 1: ")
	part1 := ReadInput()
	fmt.Print("Anna osa 2: ")
	part2 := ReadInput()
	fmt.Print("Anna osa 3: ")
	part3 := ReadInput()
	fmt.Printf("%s-%s-%s\n", part1, part2, part3)
}
