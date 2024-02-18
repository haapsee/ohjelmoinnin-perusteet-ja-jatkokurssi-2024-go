package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna sana: ")
	input, _ := reader.ReadString('\n')
	str := strings.TrimSpace(input)

	if len(str) < 2 {
		return
	}
	if str[1] == str[len(str) - 2] {
		fmt.Printf("Toinen ja toiseksi viimeinen kirjain on %s\n", string(str[1]))
		return
	}
	fmt.Printf("Toinen ja toiseksi viimeinen kirjain eroavat\n")


}
