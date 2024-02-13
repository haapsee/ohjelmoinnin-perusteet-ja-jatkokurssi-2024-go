package main

import (
	"fmt"
)

func main() {
	fmt.Print("Anna nimesi: ")
	var name string
	fmt.Scanln(&name)
	fmt.Printf("!%s!%s!\n", name, name)
}
