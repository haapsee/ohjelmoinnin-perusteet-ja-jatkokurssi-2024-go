package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadInput(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna merkkijono: ")
	str := ReadInput(reader)
	fmt.Print("Anna määrä: ")
	count, _ := strconv.ParseInt(ReadInput(reader), 10, 64)
	newString := ""
	for i := 0; i < int(count); i++ {
		newString = fmt.Sprintf("%s%s", newString, str)
	}
	fmt.Println(newString)
}
