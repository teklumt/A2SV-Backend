package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Print("Enter the string: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

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
