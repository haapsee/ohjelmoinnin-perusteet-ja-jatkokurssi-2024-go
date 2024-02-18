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
	fmt.Print("Leveys: ")
	input, _ := reader.ReadString('\n')
	width, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	for i := 0; i < int(width); i++ {
		fmt.Print("#")
	}
	fmt.Println("")
}
