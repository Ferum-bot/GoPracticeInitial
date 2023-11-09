package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice []int32
	fmt.Println("Length: ", len(slice))

	if slice == nil {
		fmt.Printf("Slice is nil\n")
	}

	secondSlice := []int32{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("Second slice == %#v\n", secondSlice)

	slicedSlice := secondSlice[1:4]
	fmt.Printf("Sliced slice == %#v\n", slicedSlice)

	secondSlice = append(secondSlice, 8)
	fmt.Printf("Second slice(appended) == %#v\n", secondSlice)

	var customSlice []int32
	//customSlice := make([]int32, 0, 1_000) Single allocation
	for i := 0; i < 1_000; i++ {
		customSlice = appendInt(customSlice, int32(i))
	}
	fmt.Printf("Custom slice: length(%d), capacity(%d) \n", len(customSlice), cap(customSlice))

	firstStringSlice := []string{"A", "B"}
	secondStringSlice := []string{"C", "D", "E"}
	fmt.Println(concatStrings(firstStringSlice, secondStringSlice))

	floatSlice := []float64{2.0, 1.0, 3.4, 4}
	fmt.Println(medium(floatSlice))
}

func appendInt(intSlice []int32, value int32) []int32 {
	insertIndex := len(intSlice)

	if len(intSlice) < cap(intSlice) {
		intSlice = intSlice[:len(intSlice)+1]
	} else {
		newLength := 2*len(intSlice) + 1
		fmt.Printf("Reallocating slice for %d \n", newLength)
		copySlice := make([]int32, newLength)
		copy(copySlice, intSlice)
		intSlice = copySlice[:len(intSlice)+1]
	}

	intSlice[insertIndex] = value
	return intSlice
}

func concatStrings(firstString, secondString []string) []string {
	resultLength := len(firstString) + len(secondString)
	resultString := make([]string, resultLength)

	copy(resultString, firstString)
	copy(resultString[len(firstString):], secondString)

	return resultString
}

func medium(values []float64) float64 {
	sort.Float64s(values)
	targetIndex := len(values) / 2

	if len(values)%2 == 1 {
		return values[targetIndex]
	}

	targetValue := (values[targetIndex-1] + values[targetIndex]) / 2
	return targetValue
}
