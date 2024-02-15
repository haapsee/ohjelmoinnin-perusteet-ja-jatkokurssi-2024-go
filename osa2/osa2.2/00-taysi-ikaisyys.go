package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	fmt.Print("Kuinka vanha olet? ")
	num, _ := strconv.ParseInt(ReadInput(), 10, 64)
	if (num < 18) {
		fmt.Println("Et ole täysi-ikäinen!")
	} else {
		fmt.Println("Olet täysi-ikäinen")
	}
}
