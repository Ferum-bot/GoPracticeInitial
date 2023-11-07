package main

import (
	"fmt"
	"unicode/utf8"
)

func banner(text string, width int) {
	padding := (width - len(text)) / 2
	for i := 0; i < padding; i++ {
		fmt.Print(" ")
	}
	fmt.Println(text)
	for i := 0; i < width; i++ {
		fmt.Print("-")
	}
	fmt.Println()
}

func isPalindrome(input string) bool {
	leftIndex := 0
	rightIndex := utf8.RuneCountInString(input) - 1

	for leftIndex <= rightIndex {
		if input[leftIndex] == input[rightIndex] {
			leftIndex++
			rightIndex--
		} else {
			return false
		}
	}

	return true
}

func main() {
	banner("Go", 6)

	fmt.Println(isPalindrome("g"))
	fmt.Println(isPalindrome("go"))
	fmt.Println(isPalindrome("gog"))
	fmt.Println(isPalindrome("gogo"))
}
