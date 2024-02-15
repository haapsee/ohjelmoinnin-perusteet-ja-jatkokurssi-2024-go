package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func ReadPIN(msg string) func() string {
	reader := bufio.NewReader(os.Stdin)

	return func() string {
		fmt.Print(msg)
		input, _ := reader.ReadString('\n')
		return strings.TrimSpace(input)
	}
}

func main() {
	reader := ReadPIN("PIN-koodi: ")
	i := 1

	for true {
		if reader() == "4321" {
			break
		}
		fmt.Println("Väärin")
		i++
	}
	if i == 1 {
		fmt.Println("Oikein, tarvitsit vain yhden yrityksen!")
		return
	}
	fmt.Printf("Oikein tarvitsit %d yritystä\n", i)
}
