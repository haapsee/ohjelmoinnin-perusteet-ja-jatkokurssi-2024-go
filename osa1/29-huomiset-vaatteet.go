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
	fmt.Println("Kerro huominen sääennuste: ")
	fmt.Print("Lämpötila: ")
	num1, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Sataako (kyllä/ei): ")
	rains := ReadInput() == "kyllä"
	fmt.Println("Pue housut ja t-paita")

	// Ota myös pitkähihainen paita
	// Pue päälle takki
	// Suosittelen lämmintä takkia
	// Kannattaa ottaa myös hanskat
	// Muista sateenvarjo!
	if (num1 <= 20) {
		fmt.Println("Ota myös pitkähihainen paita")
	}
	if (num1 <= 10) {
		fmt.Println("Pue päälle takki")
	}
	if (num1 <= 5) {
		fmt.Println("Kannattaa ottaa myös hanskat")
	}
	if (rains) {
		fmt.Println("Muista sateenvarjo!")
	}
}
