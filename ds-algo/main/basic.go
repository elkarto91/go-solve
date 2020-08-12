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

}
