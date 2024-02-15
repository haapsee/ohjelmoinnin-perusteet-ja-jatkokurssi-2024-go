package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

type Person struct {
	Name string
	Age int8
}

func ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func MinMax(x Person, y Person) Person {
	if (x.Age > y.Age) {
		return x
	}
	return y
}

func CreatePersonFromInput() Person {
	fmt.Print("Nimi: ")
	name := ReadInput()
	fmt.Print("Ikä: ")
	age, _ := strconv.ParseInt(ReadInput(), 10, 8)
	return Person{name, int8(age)}
}

func main() {
	fmt.Println("Henkilö 1:")
	x := CreatePersonFromInput()
	fmt.Println("Henkilö 2:")
	y := CreatePersonFromInput()

	if (x.Age == y.Age) {
		fmt.Printf("%s ja %s ovat yhtä vanhoja!\n", x.Name, y.Name)
	} else {
		fmt.Printf("Vanhempi on: %s\n", MinMax(x, y).Name)
	}
}
