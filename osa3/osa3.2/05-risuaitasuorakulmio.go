package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadNumber(r *bufio.Reader) int {
	input, _ := r.ReadString('\n')
	num, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	return int(num)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Leveys: ")
	width := ReadNumber(reader)
	fmt.Print("Korkeus: ")
	height := ReadNumber(reader)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print("#")
		}
		fmt.Println("")
	}

}
