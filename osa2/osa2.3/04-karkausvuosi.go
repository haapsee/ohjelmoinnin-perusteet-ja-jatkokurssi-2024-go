package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func CheckLeapYear(y int) bool {
	return !(y % 4 != 0 || (y % 100 == 0 && y % 400 != 0))
}

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Anna vuosi: ")
	input, _ := r.ReadString('\n')
	year, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if CheckLeapYear(int(year)) {
		fmt.Println("Vuosi on karkausvuosi.")
	} else {
		fmt.Println("Vuosi ei ole karkausvuosi.")
	}
}
