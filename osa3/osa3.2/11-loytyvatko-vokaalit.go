package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("anna merkkijono: ")
	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)

	foundStr := "%s löytyy"
	notFoundStr := "%s ei löydy"

	a := fmt.Sprintf(notFoundStr, "a")
	if strings.Contains(input, "a") {
		a = fmt.Sprintf(foundStr, "a")
	}

	e := fmt.Sprintf(notFoundStr, "e")
	if strings.Contains(input, "e") {
		e = fmt.Sprintf(foundStr, "e")
	}

	o := fmt.Sprintf(notFoundStr, "o")
	if strings.Contains(input, "o") {
		o = fmt.Sprintf(foundStr, "o")
	}

	fmt.Println(a)
	fmt.Println(e)
	fmt.Println(o)
}
