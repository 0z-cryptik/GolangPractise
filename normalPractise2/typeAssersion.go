package main

import "fmt"

func typeAssersion(){
	var someGreeting interface{} = "hola"

	// value returns someGreeting if it is indeed a string, else it returns the zero value of whatever type the value of someGreeting is, ok returns "true" if someGreeting is a string, else it returns false
	value, ok := someGreeting.(string)
	
	fmt.Println(value, ok)

	if ok {
		fmt.Println("it is a string")
	}
	
}