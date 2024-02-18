package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func ReadInput(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna jono 1: ")
	str1 := ReadInput(reader)
	fmt.Print("Anna jono 2: ")
	str2 := ReadInput(reader)

	if len(str1) == len(str2) {
		fmt.Println("Jonot ovat yhtä pitkät")
		return ;
	}

	if len(str2) > len(str1) {
		str1 = str2
	}
	fmt.Printf("%s on pidempi\n", str1)
}
