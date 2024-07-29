package main

import (
	"fmt"
)


func isPalindrome(word string) bool {
	
	left, right := 0, len(word)-1

	for left < right {
		if word[left] != word[right] {
			return false
		}
		left++
		right--
	}
	return true
	
}



func main() {
	var word string
	fmt.Print("Enter the string: ")
	fmt.Scan(&word)
	if isPalindrome(word) {
		fmt.Println(word + " is a palindrome.")
	} else {
		fmt.Println(word + " is not a palindrome.")
	}
}