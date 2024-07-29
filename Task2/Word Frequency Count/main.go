package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	println()
	fmt.Print("Enter the string: ")
	reader := bufio.NewReader(os.Stdin)
	inputString, _ := reader.ReadString('\n')

	frequency := make(map[string]int)

	for _, char := range inputString {
		if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') {
			upperChar := strings.ToUpper(string(char))
			frequency[upperChar]++
		}
	}

	
	fmt.Println("\nWord Frequency Count:")
	fmt.Println(strings.Repeat("-", 25))

	
	for letter, count := range frequency {
		fmt.Printf("%-10s: %d\n", letter, count)
	}

	fmt.Println(strings.Repeat("-", 25))
}
