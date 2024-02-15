package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	fmt.Print("Anna luku: ")
	num, _ := strconv.ParseInt(ReadInput(), 10, 64)

	if (num > 100) {
		fmt.Println("Luku oli suurempi, kuin 100")
		num -= 100
		fmt.Println("Nyt luvun arvo on pienentynyt sadalla")
		fmt.Printf("Arvo on nyt %d\n", num)
	}

	fmt.Printf("%d taitaa olla onnenlukuni!\n", num)
	fmt.Println("Hyvää päivänjatkoa!")
}
