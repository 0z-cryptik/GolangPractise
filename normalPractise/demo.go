package main

import (
	"fmt"
	"math"
)

const twenty int = 20

func main(){
	myMap := make(map[string]int)

	//fmt.Println(myMap)

	myMap["one"] = 1
	myMap["two"] = 2

	//fmt.Println(myMap)

	fmt.Println(add(2, 5))
	fmt.Println(random(12))

	var one, two, three int = 1,2,3
	fmt.Println(one, two, three, twenty)

	var num int = 0

	// counts from 0 - 9
	for i := 0; i < 10; i++ {
		num += 1
	}

	fmt.Println(num)

	var anotherNum int = 1

	// like while in JS
	for anotherNum < 500 {
		anotherNum += anotherNum
	}

	fmt.Println(anotherNum)
	fmt.Println(raisePow(3, 2, 10), raisePow(3, 3, 20))
}

func add(x int, y int) int {
	return x + y
}

func random(someNumber int) (x, y int) {
	x = someNumber/2
	y = someNumber + 5
	return
}

func raisePow(a, b, limit float64) float64 {
	if v := math.Pow(a, b); v < limit {
		return v
	} 
	return limit
}