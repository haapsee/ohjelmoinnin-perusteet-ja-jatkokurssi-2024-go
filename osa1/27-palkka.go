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
	// Tuntipalkka: 8.5
	// Työtunnit: 3
	// Viikonpäivä: maanantai
	// Palkka 25.5 euroa
	fmt.Print("Tuntipalkka: ")
	num1, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Työtunnit: ")
	num2, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Viikonpäivä: ")
	day := ReadInput()
	if (day == "sunnuntai") {
		num2 *= 2
	}
	fmt.Printf("Palkka %.2f euroa\n", num1*num2)
}
