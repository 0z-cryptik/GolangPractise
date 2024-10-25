package main

import (
	"fmt"
	"math"
	"runtime"
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
	fmt.Println(returnMachineOs())

	//same as if one == 0 ...
	switch {
	case one == 0:
		fmt.Println("one is 0")
	case one == 1:
		fmt.Println("one is 1")
	default:
		fmt.Printf("one is %v\n", one)
	}
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

func returnMachineOs() string {
	const partStatement = "Go runs on"
	os := runtime.GOOS

	switch os {
	case "darwin":
		return fmt.Sprintf("%v Darwin", partStatement)
	case "linux":
		return fmt.Sprintf("%v Linux", partStatement)
	default:
		return fmt.Sprintf("%v %v", partStatement, os)
	}
}