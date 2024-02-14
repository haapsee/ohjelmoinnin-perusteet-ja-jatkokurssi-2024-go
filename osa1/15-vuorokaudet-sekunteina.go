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
	fmt.Print("Kuinka monen vuorokauden sekunnit tulostetaan? ")
	days, _ := strconv.Atoi(ReadInput())
	fmt.Println(days*24*60*60)
}
