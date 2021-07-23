package main

import "fmt"

func main() {

	var array [4]int
	array[0] = 111
	array[1] = 222
	array[2] = 666
	array[3] = 999
	fmt.Println(array)
	fmt.Println("len of array",len(array))

	for i:= 1; i < len(array); i++{
		array[i] = 0
	}

	fmt.Println(array)

	for index, _ := range array{
		array[index]++
		//fmt.Printf("array[%d] = %d\n", index, value)
	}
	fmt.Println(array)
}
