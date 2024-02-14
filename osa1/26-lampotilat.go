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
	fmt.Print("Anna lämpötila (F): ")
	fahrenheits, _ := strconv.ParseFloat(ReadInput(), 64)
	//  (50°F − 32) / 1,8 = 10°C
	celsius := (fahrenheits - 32.0) / 1.8
	// 101 fahrenheit-astetta on 38.333333333333336 celsius-astetta
	fmt.Printf("%g fahrenheit-astetta on %.2f celsius-astetta\n", fahrenheits, celsius)
	if (celsius < 0) {
		fmt.Println("Paleltaa!")
	}
}
