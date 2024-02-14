package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func main() {
	fmt.Print("Luku 1: ")
	num1, _ := strconv.ParseInt(ReadInput(), 10, 64)
	fmt.Print("Luku 2: ")
	num2, _ := strconv.ParseInt(ReadInput(), 10, 64)
	fmt.Print("Komento: ")
	command := ReadInput()
	if (command == "summa") {
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	}
	if (command == "tulo") {
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	}
	if (command == "erotus") {
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	}
}


