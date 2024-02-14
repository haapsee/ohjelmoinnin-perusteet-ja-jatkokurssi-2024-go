package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func ReadInput() int64 {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	res, _ := strconv.ParseInt(strings.TrimSpace(result), 10, 64)
	return res
}

func main() {
	fmt.Print("Anna luku: ")
	num := ReadInput()
	if (num < 1000) {
		fmt.Println("Luku on pienempi kuin 1000")
	}
	if (num < 100) {
		fmt.Println("Luku on pienempi kuin 100")
	}
	if (num < 10) {
		fmt.Println("Luku on pienempi kuin 10")
	}
	fmt.Println("Kiitos!")
}


