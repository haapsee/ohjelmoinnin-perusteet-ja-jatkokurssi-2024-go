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
	fmt.Print("Mikä on nimesi? ")
	name := ReadInput()
	if (name != "Jerry") {
		fmt.Print("Kuinka monta annosta keittoa? ")
		amountOfSoups, _ := strconv.ParseInt(ReadInput(), 10, 64)
		fmt.Printf("Kokonaishinta on %.2f\n", (float64(amountOfSoups) * 5.9))
	}
	fmt.Println("Seuraava!")
}
