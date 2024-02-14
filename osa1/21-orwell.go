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
	fmt.Print("Anna luku: ")
	num, _ := strconv.ParseInt(ReadInput(), 10, 64)
	if (num == 1984) {
		fmt.Println("Orwell")
	}
}
