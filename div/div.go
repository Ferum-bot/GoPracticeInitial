package main

import "fmt"

func div(a, b int) int {
	return a / b
}

func saveDiv(a, b int) (res int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Error: ", e)
			err = fmt.Errorf("%v", e)
		}
	}()
	return a / b, nil
}

func main() {
	fmt.Println(saveDiv(1, 0))
}
