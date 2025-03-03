package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// error is printed with this format
func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return MyError{
		time.Now(), 
		"it didn't work",
	}
}