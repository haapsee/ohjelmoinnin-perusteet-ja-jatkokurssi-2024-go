package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func main() {
	fmt.Print("Anna luku 1: ")
	num1, _ := strconv.Atoi(ReadInput())
	fmt.Print("Anna luku 2: ")
	num2, _ := strconv.Atoi(ReadInput())
	fmt.Printf("Lukujen summa %d\n", (num1+num2))
	fmt.Printf("Lukujen tulo %d\n", (num1*num2))
}
