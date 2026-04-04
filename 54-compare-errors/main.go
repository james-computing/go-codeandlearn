package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Message string
	Status  int
}

func (e CustomError) Error() string {
	return e.Message
}

var ErrFirstError CustomError = CustomError{Message: "first error"}

func someFunction() error {
	//return ErrFirstError
	return fmt.Errorf("some function err: %w", ErrFirstError)
}

func anotherFunction() error {
	return CustomError{Message: "original error: something went wrong", Status: 400}
}

func main() {
	// errors.Is
	err := someFunction()
	if err != nil {
		if errors.Is(err, ErrFirstError) {
			fmt.Println("sentinel error found")
		}
	}

	// errors.As
	var customError CustomError
	err = anotherFunction()
	if errors.As(err, &customError) {
		fmt.Println("Extrated CustomError:", customError)
	}
}
