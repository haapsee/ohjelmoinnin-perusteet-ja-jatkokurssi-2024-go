package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func ReadNumber(r *bufio.Reader) float64 {
	input, _ := r.ReadString('\n')
	value, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)
	return float64(value)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := []float64{}

	fmt.Println("Syötä kokonaislukuja, 0 lopettaa: ")

	for true {
		fmt.Print("Luku: ")
		num := ReadNumber(reader)
		if num == 0 {
			break
		}
		nums = append(nums, num)
	}

	numsSum := 0.0
	pos := 0
	neg := 0

	for i := 0; i < len(nums); i++ {
		numsSum += nums[i]
		if (nums[i] > 0) {
			pos++
		} else {
			neg++
		}
	}

	fmt.Printf("Lukuja yhteensä: %d\n", len(nums))
	fmt.Printf("Lukujen summa: %.2f\n", numsSum)
	fmt.Printf("Lukujen keskiarvo: %.2f\n", numsSum / float64(len(nums)))
	fmt.Printf("Positiivisia: %d\n", pos)
	fmt.Printf("Negatiivisia: %d\n", neg)

}
