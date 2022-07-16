package main

import (
	"fmt"
	"reflect"
)

const nameOfInstitution = "Alt Africa"

var myName string = "alt africa"

func main() {
	myName = ""
	var name string = "alt africa"
	fmt.Println(name)
	var age int = 50
	fmt.Println(age)

	var price float64 = 5.64
	fmt.Println(price)

	var decision bool
	fmt.Println(decision)
	decision = true
	fmt.Println(decision)
	fmt.Println(nameOfInstitution)

	var value int64 = 50
	var newVariable int32

	newVariable = int32(value)
	fmt.Println(newVariable)

	fmt.Printf("checking the type of newVariable %T", newVariable)

	// %v, %d, %s

	var myAge string = "hello"
	fmt.Println()
	fmt.Printf("MY AGE IS %v and i have the type of %T", myAge, myAge)
	fmt.Println()
	//reflect
	fmt.Println(reflect.TypeOf(decision))

	a := 6.000
	fmt.Printf("type is %T", a)

	A, B, C := "NEW", 12, 12
	println()
	fmt.Println(A, B, C)
	fmt.Printf("the types of A, B, C is: %T %T %T", A, B, C)

}

func testing() {
	var ab int
	ab = 80
	myName = "hello"
	fmt.Println(ab)
}
