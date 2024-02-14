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
	fmt.Print("Montako kertaa viikossa syöt Unicafessa? ")
	num1, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Unicafe-lounaan hinta? ")
	num2, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Paljonko käytät viikossa ruokaostoksiin? ")
	num3, _ := strconv.ParseFloat(ReadInput(), 64)

	fmt.Printf("Kustannukset keskimäärin:\nPäivässä %g euroa\nViikossa %g euroa\n", ((num1*num2+num3)/7), (num1*num2+num3))
	// Kustannukset keskimäärin:
	// Päivässä 5.5 euroa
	// Viikossa 38.5 euroa
}
