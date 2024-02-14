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
	fmt.Print("Luku 1: ")
	num1, _ := strconv.Atoi(ReadInput())
	fmt.Print("Luku 2: ")
	num2, _ := strconv.Atoi(ReadInput())
	fmt.Print("Luku 3: ")
	num3, _ := strconv.Atoi(ReadInput())
	fmt.Print("Luku 4: ")
	num4, _ := strconv.Atoi(ReadInput())
	sumOfNums := num1+num2+num3+num4
	fmt.Printf("Lukujen summa on %d ", sumOfNums)
	fmt.Printf("ja keskiarvo %g\n", (float64(sumOfNums)/4))
}
