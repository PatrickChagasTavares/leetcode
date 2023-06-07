package main

import "fmt"

const (
	first  = iota + 6
	second = 2 << iota
	three  = 32 >> 5
)

func main() {

	slice := make([]int, 0, 10)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))

	slice = append(slice, 1)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 2)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 3)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 4)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 5)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 6)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 7)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 8)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 9)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 10)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))
	slice = append(slice, 10)
	fmt.Printf("value: %v, cp: %v, len: %v \n", slice, cap(slice), len(slice))

	fmt.Println("first ", first, second, three)
	// fmt.Println("second ", second)

}
