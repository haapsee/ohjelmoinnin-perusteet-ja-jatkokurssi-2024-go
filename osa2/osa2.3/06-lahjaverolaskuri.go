package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Lahjan suuruus: ")
	input, _ := reader.ReadString('\n')
	value, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	res := float64(0)

	if value < 5000 {
		res = float64(0)
	} else if value < 25000 {
		res = float64(value - 5000) * 0.08 + 100
	} else if value < 55000 {
		res = float64(value - 25000) * 0.1 + 1700
	} else if value < 200000 {
		res = float64(value - 5500) * 0.12 + 4700
	} else if value < 1000000 {
		res = float64(value - 200000) * 0.15 + 22100
	} else {
		res = float64(value - 1000000) * 0.17 + 142100
	}

	if res == 0 {
		fmt.Println("Ei veroa!")
	} else {
		fmt.Printf("Vero: %.2f euroa\n", res)
	}

}
