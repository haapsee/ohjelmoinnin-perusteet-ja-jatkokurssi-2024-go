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
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func main() {
	fmt.Print("Anna luku: ")
	number, _ := strconv.Atoi(ReadInput())
	fmt.Printf("Kun kerrotaan %d luvulla 5, saadaan %d\n", number, (number*5))
}
