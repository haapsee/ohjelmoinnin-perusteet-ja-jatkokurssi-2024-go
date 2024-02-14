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
	fmt.Print("Anna nimi: ")
	name := ReadInput()
	fmt.Print("Anna vuosi: ")
	year := ReadInput()
	fmt.Printf("%s on urhea ritari, syntynyt vuonna %s. Eräänä aamuna %s heräsi kovaan meluun: lohikäärme lähestyi kylää. Vain %s voisi pelastaa kylän asukkaat.\n", name, year, name, name)
}
