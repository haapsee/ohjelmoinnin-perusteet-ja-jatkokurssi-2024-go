package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadInput(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	value, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	return int(value)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna pisteet [0-100]: ")
	points := ReadInput(reader)

	res := ""
	if points < 0 || points > 100 {
		res = "mahdotonta!"
	} else if points < 50 {
		res = "hylätty"
	} else if points < 60 {
		res = "1"
	} else if points < 70 {
		res = "2"
	} else if points < 80 {
		res = "3"
	} else if points < 90 {
		res = "4"
	} else {
		res = "5"
	}

	fmt.Printf("Arvosana: %s\n", res)
}
