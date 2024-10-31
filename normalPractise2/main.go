package main

import (
	"fmt"
)

// var someNums = []int {1,2,3,4,5,6,7,8,9}

func main(){
	fmt.Println("we up!!!")

	// for index, value := range someNums {
	// 	fmt.Println(index, value)
	// }

	myMap := map[string]string {
		"name": "Ronaldo",
		"club": "Al Nassr",
		"age": "39",
		"height": "6'2",
	}

	myMap["bestie"] = "Adebayor"

	fmt.Println(myMap, myMap["name"])

	_, present := myMap["width"]

	fmt.Println(present)
	fmt.Println(v.Multiply())
	typeAssersion()
	typeSwitches()
}