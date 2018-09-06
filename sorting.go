package main

import (
	"sort"
	"fmt"
)

func main() {
	// Sort slice of strings
	stringValues := []string{"banana", "cat", "apple"}
	sort.Strings(stringValues)
	fmt.Println("stringValues", stringValues)
	// Output: ["apple", "banana", "cat"]

	// Sort slice of int
	intValues := []int{10, 2, 5}
	sort.Ints(intValues)
	fmt.Println("intValues", intValues)
	// Output: [2 5 10]

	// Sort slice of float64
	float64Values := []float64{4.5, 3.4, 1.2}
	sort.Float64s(float64Values)
	fmt.Println("float64Values", float64Values)
	// Output: [1.2 3.4 4.5]

	// Sort slice of int
	int64Values := []int64{17, 8, 10}
	int64AsIntValues := make([]int, len(int64Values))
	for i, val := range int64Values {
		int64AsIntValues[i] = int(val)
	}
	sort.Ints(int64AsIntValues)
	for i, val := range int64AsIntValues {
		int64Values[i] = int64(val)
	}
	fmt.Println("int64Values", int64Values)
	// Output: [8, 10, 17]

	// Sort slice of float32
	float32Values := []float32{7.9, 2.4, 5.7}
	float32AsFloat64Values := make([]float64, len(float32Values))
	for i, val := range float32Values {
		float32AsFloat64Values[i] = float64(val)
	}
	sort.Float64s(float32AsFloat64Values)
	for i, val := range float32AsFloat64Values {
		float32Values[i] = float32(val)
	}
	fmt.Println("float32Values", float32Values)
	// Output: [2.4 5.7 7.9]
}
