package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
)

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	result, _ := reader.ReadString('\n')
	return strings.TrimSpace(result)
}

func main() {
	fmt.Print("Etunimi: ")
	firstname := ReadInput()
	fmt.Print("Sukunimi: ")
	lastname := ReadInput()
	fmt.Print("Katuosoite: ")
	streetAddress := ReadInput()
	fmt.Print("Postinumero ja kaupunki: ")
	postalcodeAndCity := ReadInput()
	fmt.Printf("%s %s\n%s\n%s\n", firstname, lastname, streetAddress, postalcodeAndCity)
}
