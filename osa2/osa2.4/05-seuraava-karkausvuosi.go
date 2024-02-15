package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadYear(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	year, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	return int(year)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Vuosi: ")
	startYear := ReadYear(reader)
	nextLeapYear := startYear + 1
	for nextLeapYear % 4 != 0 || (nextLeapYear % 100 == 0 && nextLeapYear % 400 != 0) {
		nextLeapYear++
	}
	fmt.Printf("Vuotta %d seuraava karkausvuosi on %d\n", startYear, nextLeapYear)
}
