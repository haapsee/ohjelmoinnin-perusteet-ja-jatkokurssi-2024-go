package main

import (
	"fmt"
	"os"
	"bufio"
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
	fmt.Print("Anna a: ")
	a, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Anna b: ")
	b, _ := strconv.ParseFloat(ReadInput(), 64)
	fmt.Print("Anna c: ")
	c, _ := strconv.ParseFloat(ReadInput(), 64)

	x := (-b + math.Sqrt(math.Pow(b, 2) - 4 * a * c))/(2 * a)
	y := (-b - math.Sqrt(math.Pow(b, 2) - 4 * a * c))/(2 * a)

	fmt.Printf("Juuret ovat %.2f ja %.2f\n", x, y)
}
