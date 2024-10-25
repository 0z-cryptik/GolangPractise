package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

//returns string and error (if there's no name provided)
func Hello(name string) (string, error){
	if name == "" {
		return "", errors.New("empty name")
	}

	if name == "Eniitan" {
		message := fmt.Sprintf("%v, you can suck my ...", name)
		return message, nil
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error){
	//map is like a dict in python, map[string]string means both key and value are of type string
	messages := make(map[string]string)
	// Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
        // the name.
		messages[name] = message
	}

	return messages, nil
}

func randomFormat() string {
	formats := [] string {
		"Hi, %v. Welcome!",
		"Great to see you %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}