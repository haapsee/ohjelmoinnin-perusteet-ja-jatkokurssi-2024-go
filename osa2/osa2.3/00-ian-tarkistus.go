package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadInput(r *bufio.Reader) int {
	input,  _ := r.ReadString('\n')
	value, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	return int(value)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Kerro ikäsi? ")
	age := ReadInput(reader)

	if age >= 0 && age < 5 {
		fmt.Println("En usko, että osaat kirjoittaa...")
	} else if age < 0 || age > 150 {
		fmt.Println("Taitaa olla virhe")
	} else {
		fmt.Printf("OK, olet siis %d-vuotias\n", age)
	}
}
