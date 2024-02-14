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
	fmt.Print("Kuinka paljon pisteitä? ")
	points, _ := strconv.ParseFloat(ReadInput(), 64)
	if (points < 100) {
		points *= 1.1
		fmt.Println("Sait 10% bonusta")
	} else {
		points *= 1.15
		fmt.Println("Sait 15% bonusta")
	}
	fmt.Printf("Pisteitä on nyt %.2f\n", points)
}
