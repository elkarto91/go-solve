package main

import (
	"fmt"
	"time"
)

//Constant Declaration
const MAX = 9999999

//Global Variable
var i int32

func main() {

	fmt.Println(" Constant Max ", MAX)

	//Type Conversion
	fmt.Println("max value in string ", string(MAX))

	//Regular For Loop
	for i = 0; i < MAX; i++ {

		if i < 10 {
			fmt.Println("I value is", i)
		}
	}

	//While Loop Implementation
	for i < 30 {
		fmt.Println("i value", i)
		i++
	}

	//Switch
	switch i {

	case 10:
		fmt.Println("10")
	case 20:
		fmt.Println("20")
	default:
		fmt.Println("Not in the list ")

	}

	//Time Function In Switch
	switch time.Now().Weekday() {

	case time.Sunday:
		fmt.Println("Its Sun")
	case time.Monday:
		fmt.Println("Its Mon")
	case time.Tuesday:
		fmt.Println("Its Tue")
	case time.Wednesday:
		fmt.Println("Its Wed")
	case time.Thursday:
		fmt.Println("Its Thur")
	case time.Friday:
		fmt.Println("Its Fri")
	case time.Saturday:
		fmt.Println("Its Sat")
	default:
		fmt.Println("Not in list")

	}

	//Arrays
	//Declaration
	var intArr []int32
	var stringArr []string

	//Intialization
	stringArr = []string{"First", "Second"}

	for _, v := range stringArr {
		fmt.Println(v)
	}

	intArr = []int32{10, 200, 3000}
	for _, u := range intArr {

		fmt.Println(u)
	}

	//Two dimensional arrays
	var twod [][]int32

	twod = [][]int32{{1, 2}, {3, 4}, {5, 6}, {7, 8}}

	//For loop with range
	for _, k := range twod {

		fmt.Println(k)
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {

			fmt.Println("Sum Values Matrix")
			twod[i][j] = int32(i + j)
		}
	}
	fmt.Println("Updated Matrix", twod)

	//Slices

	slic := make([]string, 3)
	slic = []string{"Henry", "Ozil", "Ramsey", "Mertesacker", "Eduardo", "Walcott"}
	fmt.Println("Slice Data", slic)

	slic[0] = "Aubameyang"
	fmt.Println("New Slice Value", slic)

	slic = append(slic, "Arteta")
	fmt.Println("Appended Slice", slic)

	copiedslice := make([]string, 5)
	fmt.Println("Copied Slice before copy", copiedslice)

	copy(copiedslice, slic)
	fmt.Println("Copied SLice after copy", copiedslice)

	fmt.Println("Slice range 2-4 ", slic[2:4])
	fmt.Println("slice range upto 3", slic[:3])
	fmt.Println("slice range from 4", slic[1:])

	//Map key value pair

	mapData := make(map[string]int)
	fmt.Println("New Map", mapData)

	mapData["Ozil"] = 350000
	mapData["Kane"] = 000000

	fmt.Println("Updated Map", mapData)

	fmt.Println("Length of Map", len(mapData))

	delete(mapData, "Kane")

	fmt.Println("Updated Map after deletion", mapData)

	fmt.Println(" Calling Function")

	f, g := functionForExecution("try me")
	fmt.Println("Returned values are : ", f, g)

	sum := nNoOfArgumentsFunction(2, 45, 12, 25, 34, 3, 242)
	fmt.Println("Sum", sum)

	sum = nNoOfArgumentsFunction(2, 45)
	fmt.Println("Sum", sum)

	fmt.Println("Closure Funtion Example")
	returnedValue := closureFunctionExample()
	fmt.Println("CLOSURE DONE", returnedValue())

	returnedValue = closureFunctionExample()
	fmt.Println("CLOSURE DONE", returnedValue())

	fmt.Println("Factorial ")
	factorialResult0 := factorial(0)
	fmt.Println("Factorial 0 ", factorialResult0)

	factorialResult1 := factorial(1)
	fmt.Println("Factorial 1 ", factorialResult1)

	factorialResult9 := factorial(9)
	fmt.Println("Factorial 9 ", factorialResult9)

	fmt.Println("Fibonacci ")
	fibonacciResult0 := fibonacci(0)
	fmt.Println("Fibonacci 0 ", fibonacciResult0)

	fibonacciResult1 := fibonacci(1)
	fmt.Println("Fibonacci 1 ", fibonacciResult1)

	fibonacciResult9 := fibonacci(9)
	fmt.Println("Fibonacci 9 ", fibonacciResult9)

	fmt.Println("Fibonacci result for 7 ")
	for m := 0; m <= 7; m++ {

		fmt.Println(fibonacci(m))
	}

}
func functionForExecution(string2 string) (int, int) {

	fmt.Println(" The string is", string2)
	return 24, 11

}

func nNoOfArgumentsFunction(arg ...int) int {

	sum := 0
	for _, v := range arg {
		sum = sum + v
	}
	return sum
}

func closureFunctionExample() func() int {
	fmt.Println(" The function houses another function which is declared anonymously...as in without a name")
	//This part is body of the outer function, the body of inner function is housed inside it as follows.
	//The closure function returns an integer
	i := -0
	return func() int {
		i++
		switch i {
		case 1:
			fmt.Println("i is one")
		case 2:
			fmt.Println(" i is two")
		case 3:
			fmt.Println("i is three")
		default:
			fmt.Println(" i isnt the same ")
		}
		return i
	}

	//Outer function doesnt need a return statement
}

func factorial(num int) int {

	if num == 0 {
		return 1
	}
	return num * (factorial(num - 1))

}
func fibonacci(num int) int {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1

	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}
