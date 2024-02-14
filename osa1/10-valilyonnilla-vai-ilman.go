package main

import (
	"fmt"
)

func main() {
	name := "Teppo Testaaja"
	age := 20
	skill1 := "python"
	skill2 := "java"
	skill3 := "programming"
	skillLevel1 := "beginner"
	skillLevel2 := "veteran"
	skillLevel3 := "semi professional"
	minimumWage := 2000
	maximumWage := 3000

	fmt.Printf("Nimeni on %s, olen %d-vuotias\n\n", name, age)
	fmt.Printf("Taitoihini kuuluvat\n")
	fmt.Printf(" - %s (%s)\n", skill1, skillLevel1)
	fmt.Printf(" - %s (%s)\n", skill2, skillLevel2)
	fmt.Printf(" - %s (%s)\n\n", skill3, skillLevel3)
	fmt.Printf("Haen työtä, josta maksetaan palkkaa %d-%d euroa kuussa\n", minimumWage, maximumWage)
}
