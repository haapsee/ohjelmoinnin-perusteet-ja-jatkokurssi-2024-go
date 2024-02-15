package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"slices"
)

func ReadInput(r *bufio.Reader) string {
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Anna nimesi: ")
	name := ReadInput(reader)

	if slices.Contains([]string{"Tupu", "Hupu", "Lupu"}, name) {
		fmt.Println("Olet luultavasti Aku Ankan veljenpoika.")
	} else if slices.Contains([]string{"Mortti", "Vertti"}, name) {
		fmt.Println("Olet luultavasti Mikki Hiiren veljenpoika.")
	} else {
		fmt.Println("Et ole kenenkään tuntemani hahmon veljenpoika.")
	}
}
