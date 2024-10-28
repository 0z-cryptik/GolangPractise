package main

import (
	"fmt"
)

type MyStruct struct {
	X int
	Y int
}

type Mystruct2 struct {
	name string
	school string
}

func main(){
	var obj MyStruct = MyStruct{}
	obj.X, obj.Y = 20, 619
	fmt.Println(obj)
	fmt.Println(obj.X, obj.Y)

	var obj2 Mystruct2 = Mystruct2{
		name: "Bojan", 
		school: "La Masia",
	}

	fmt.Println(obj2)

	var arrayOfTwoInts [2]int
	arrayOfTwoInts[0] = 12
	arrayOfTwoInts[1] = 46
	fmt.Println(arrayOfTwoInts)

	arrayOfFiveStrings := [5]string {
		"one", "two", "three", "four", "five",
	}

	//changes made to a slice affects the original array
	var sliceOfFourStrings []string =  arrayOfFiveStrings[0:4]
	mySlice := []int {1,2,4,16}

	fmt.Println(arrayOfFiveStrings, sliceOfFourStrings, mySlice)

	type ArrayOrSliceStruct struct {
		i int
		s string
	}

	randomArray := [4]ArrayOrSliceStruct {
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
	}

	var randomSlice []ArrayOrSliceStruct = randomArray[0:2]
	
	fmt.Println(randomArray)
	fmt.Println(randomSlice, cap(randomSlice))

	// how to create a slice with the make function
	// this creates a slice of 5 integers
	someSlice1 := make([]int, 5)

	// this creates a slice of ints with length = 0 and capacity = 5
	someSlice2 := make([]int, 0, 5)

	fmt.Printf("len of someSlice1: %d\n cap of someSlice1: %d\n", len(someSlice1), cap(someSlice1))
	fmt.Printf("len of someSlice2: %d\n cap of someSlice2: %d\n", len(someSlice2), cap(someSlice2))

	someSlice1 = append(someSlice1, 1, 2, 3)

	fmt.Println(someSlice1)


}