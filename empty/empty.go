package main

import "fmt"

func main() {
	var anyVariable any

	anyVariable = 7
	fmt.Println(anyVariable)

	anyVariable = "Hello any!"
	fmt.Println(anyVariable)

	s, ok := anyVariable.(string)
	if ok {
		fmt.Println(s)
	}

	switch anyVariable.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("Default")
	}
}

func maximum[T int | float64](numbers []T) T {
	if len(numbers) == 0 {
		return 0
	}

	maximum := numbers[0]
	for _, number := range numbers {
		if number > maximum {
			maximum = number
		}
	}

	return maximum
}
