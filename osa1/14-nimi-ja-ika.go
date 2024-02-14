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
	fmt.Print("Anna nimi: ")
	name := ReadInput()
	fmt.Print("Anna syntymävuosi: ")
	year, _ := strconv.Atoi(ReadInput())
	fmt.Printf("Moi %s, olet 30 vuotta vanha vuoden %d lopussa\n", name, (year+30))
}
