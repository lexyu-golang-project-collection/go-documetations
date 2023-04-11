package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// message := fmt.Sprintf("Hi, %v. Welcome!", name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v",
		"Hail, %v! Well met!",
	}
	println(reflect.TypeOf(formats))
	println(reflect.TypeOf(formats).String())

	return formats[rand.Intn(len(formats))]
}
