package main

import "fmt"

func typeSwitches(){
	fmt.Println(checkType("abracadabra"))
	fmt.Println(checkType(22))
	fmt.Println(checkType(false))
}

func checkType(i interface{}) string {
	switch i.(type) {
	case int:
		return "it is an int"
	case string:
		return "it is a string"
	default:
		return "no idea wtf it is"
	}
}