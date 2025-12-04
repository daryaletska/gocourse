package main

import "fmt"

func main() {
	var numbers [5]int
	fmt.Println(numbers)

	numbers[4] = 111
	fmt.Println(numbers)

	originalArray := [3]int{1, 2, 3}
	// copiedArray := originalArray

	// copiedArray[0] = 100

	fmt.Println("originalArray:", originalArray)
	// fmt.Println("copyArray:", copiedArray)

	for i := 0; i < len(originalArray); i++ {
		fmt.Println("Element at index", i, ":", originalArray[i])
	}

	for index, value := range originalArray {
		fmt.Printf("Index %d: value %d\n", index, value)
	}

	for _, __ := range originalArray {
		fmt.Printf("Value %d\n", __)
	}

	a, b := someFunction()
	fmt.Println("a:", a)
	fmt.Println("b:", b)

	length := len(originalArray)
	fmt.Println("Length of originalArray:", length)

	// comparing arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}
	fmt.Println("array1 == array2:", array1 == array2)

	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("matrix:", matrix)

	var copiedArray *[3]int = &matrix[0]
	copiedArray[0] = -1
	fmt.Println("original:", matrix[0])
	fmt.Println("copiedArray:", *copiedArray)
}

func someFunction() (int, int) {
	return 1, 2
}
