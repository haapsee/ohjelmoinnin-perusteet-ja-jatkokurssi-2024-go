package main

import (
	"os"
	"fmt"
	"sort"
	"bufio"
	"strings"
)

func ReadAndAppend(r *bufio.Reader, list []string) []string {
	input, _ := r.ReadString('\n')
	return append(list, strings.TrimSpace(input))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	list := []string{}

	fmt.Print("Anna 1. kirjain: ")
	list = ReadAndAppend(reader, list)
	fmt.Print("Anna 2. kirjain: ")
	list = ReadAndAppend(reader, list)
	fmt.Print("Anna 3. kirjain: ")
	list = ReadAndAppend(reader, list)

	sort.Strings(list)
	fmt.Printf("Keskimmäinen kirjain on %s\n", list[1])
}
