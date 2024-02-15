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
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	fmt.Print("Anna luku: ")
	input, _ := strconv.ParseFloat(ReadInput(), 64)
	num1 := int(input)
	num2 := input - float64(num1)
	fmt.Printf("Kokonaisosa: %d\nDesimaaliosa: %.2f\n", num1, num2)
}
