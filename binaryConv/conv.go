package main

import "fmt"

func main() {
	loopBin := []string{"01001001", "01101110", "01100110", "01101001",
		"01101110", "01101001", "01110100", "01100101", "00100000", "01101100",
		"01100101", "01100001", "01110010", "01101110", "01101001", "01101110",
		"01100111", "00100001"}

	for _, bin := range loopBin {
		convert(bin)
	}
}

func convert(bin string) {
	weights := []int{128, 64, 32, 16, 8, 4, 2, 1}
	sum := 0
	for i, char := range bin {
		if char == '1' {
			sum += weights[i]
		}
	}
	r := rune(sum)
	fmt.Printf("%c", r)
}