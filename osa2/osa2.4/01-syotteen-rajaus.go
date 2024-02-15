package main

import (
	"os"
	"fmt"
	"math"
	"bufio"
	"strings"
	"strconv"
)

func ReadNumber(r *bufio.Reader) float64 {
	input, _ := r.ReadString('\n')
	value, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	return value
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	num := -1.0

	for num != 0 {
		fmt.Print("Syötä luku: ")
		num = ReadNumber(reader)
		if num == 0 {
			fmt.Println("Lopetetaan...")
		} else if num < 0 {
			fmt.Println("Epäkelpo luku")
		} else {
			fmt.Printf("%.2f\n", math.Sqrt(num))
		}
	}
}
