package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
)

func ReadInput(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func find(a string, b string, index int, second bool) string {
	if len(a) < len(b) {
		return "Osajono ei esiinny merkkijonossa kahdesti."
	}

	if a[:len(b)] != b {
		return find(a[1:], b, index+1, second)
	}

	if a[:len(b)] == b {
		if second {
			return fmt.Sprintf("Osajonon toinen esiintymä on kohdassa %d.", index)
		}
		return find(a[len(b):], b, index+len(b), true)
	}
	return "Osajono ei esiinny merkkijonossa kahdesti."
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Anna merkkijono: ")
	str1 := ReadInput(reader)
	fmt.Print("Anna osajono: ")
	str2 := ReadInput(reader)

	fmt.Println(find(str1, str2, 0, false))
}
