package main

import "fmt"

type typeFn func(int) int

func main() {
	fmt.Println("Hello world")

	var numbers []int = []int{1, 2, 3, 4, 5}

	doubledNumber := transformNumber(&numbers, doubleNumber)
	tripledNumber := transformNumber(&numbers, tripleNumber)
	quadrupledNumber := transformNumber(&numbers, func(i int) int {
		return i * 4
	})

	quintupleFn := createTransformer(5)

	quintupledNumber := transformNumber(&numbers, quintupleFn)

	fmt.Println(doubledNumber)
	fmt.Println(tripledNumber)
	fmt.Println(quadrupledNumber)
	fmt.Println(quintupledNumber)
}

func transformNumber(numbers *[]int, function typeFn) []int {
	result := make([]int, len(*numbers))

	for i, v := range *numbers {
		result[i] = function(v)
	}

	return result
}

func doubleNumber(number int) int {
	return number * 2
}

func tripleNumber(number int) int {
	return number * 3
}

func createTransformer(factor int) typeFn {
	return func(number int) int {
		return number * factor
	}
}
