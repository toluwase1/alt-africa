package main

import (
	"fmt"
	"sort"
)

var firstArray [6]int

var firstSlice = []string{}

func main() {
	// firstArray[0] = 20
	// firstArray[1] = 30
	// firstArray[2] = 40
	// firstArray[3] = 80
	// firstArray[4] = 90
	// firstArray[5] = 4400
	// fmt.Println(firstArray)
	// fmt.Println(firstArray[4])
	// fmt.Println(len(firstArray))
	// fmt.Println(cap(firstArray))

	// firstSlice = append(firstSlice, "greetings")
	// firstSlice = append(firstSlice, "hello")

	// fmt.Println(firstSlice)

	// var numberArray = []int{100, 20, 50, 80, 10, 89, 800, 569}

	// fmt.Println(numberArray)

	// mySlice := numberArray[:]
	// fmt.Println(mySlice)
	// fmt.Println(numberArray[:3])
	// fmt.Println(numberArray[5:])

	// var slices []string = []string{"go", "run", "main", "slices"}
	// printIndexAndValue(slices)
	var linearSearchValues = []int{100, 20, 50, 80, 10, 89, 800, 100}
	// printFew(linearSearchValues)
	// linearSearch1(linearSearchValues)
	fmt.Println(aVeryBigSum(linearSearchValues))

	// fmt.Println(linearSearch(linearSearchValues, 10))
	fmt.Println(containsDuplicate(linearSearchValues))

}

func printIndexAndValue(slice []string) {
	// start condition, iteration
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
}

func printFew(slice []int) {
	for i := 0; i < len(slice)-9; i++ {
		fmt.Println(slice[i])
	}
}

func linearSearch1(slice []int) {
	for _, j := range slice {
		fmt.Println(j)
	}
}

func aVeryBigSum(slice []int) int {
	//var linearSearchValues = []int{100, 20, 50, 80, 10, 89, 800, 569}
	var result int
	for i := 0; i < len(slice); i++ {
		result = slice[i] + result
	}
	return result
}

func containsDuplicate(slice []int) bool {
	//var linearSearchValues = []int{100, 20, 50, 80, 10, 89, 800, 569}
	sort.Ints(slice)

	for i := 0; i < len(slice)-1; i++ {
		if slice[i] == slice[i+1] {
			return true
		}
	}
	return false
}
