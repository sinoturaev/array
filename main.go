package main

func MaxIndexOfMaxElement(numbers []int) (int , int) {
	Index := 0
	MaxElement := numbers[0]
	for index, value := range numbers{
		if MaxElement < value {
			MaxElement = value
			Index = index
		}
	}
	return Index, MaxElement
}

func IndexOfMaxElement(numbers []int) int {
	MaxElement := numbers[0]
	for _, value := range numbers{
		if MaxElement < value {
			MaxElement = value
		}
	}
	return MaxElement
}

func main() {
	IndexOfMaxElement([]int{10, 25, -5, 1, -100, 25})
}
