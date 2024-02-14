package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"math"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func main() {
	// Montako opiskelijaa? 8
	// Mikä on ryhmän koko? 4
	// Ryhmien määrä: 2
	fmt.Print("Montako opiskelijaa? ")
	students, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Mikä on ryhmän koko? ")
	groupSize, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Printf("Ryhmien määrä: %g\n", math.Ceil(students/groupSize))
}
