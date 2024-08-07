package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func CountCharacterFrequency(input string) map[rune]int {
	frequency := make(map[rune]int)

	reg := regexp.MustCompile(`[\p{P}-[._]]`)
	input = reg.ReplaceAllString(input, "")

	input = strings.ToUpper(input)

	for _, char := range input {
		if char != ' ' {
			frequency[char]++
		}
	}

	return frequency
}

func main() {
	var input string

	fmt.Print("Enter the string: ")
	fmt.Scanln(&input)

	frequency := CountCharacterFrequency(input)

	var keys []rune
	for k := range frequency {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	fmt.Println("Character frequencies:")
	for _, char := range keys {
		fmt.Printf("%c: %d\n", char, frequency[char])
	}
}
