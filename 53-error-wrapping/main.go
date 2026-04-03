package main

import (
	"errors"
	"fmt"
)

func main() {
	err := secondFunction()
	fmt.Println("original error:", err)

	err = SomeFunction()

	// unwrap error
	innerError := errors.Unwrap(err)
	fmt.Println("Inner error:", innerError)
}

func firstFunction() error {
	return errors.New("original error: something went wrong in firstFunction")
}

func secondFunction() error {
	firstErr := firstFunction()
	if firstErr != nil {
		// wrap error
		return fmt.Errorf("failed in first function: %w", firstErr)

		// join
		//secondErr := errors.New("failed in second function")
		//return errors.Join(firstErr, secondErr)
	}
	return nil
}

type CustomError struct {
	Message string
	Wrapped error
}

func (e CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Wrapped)
}

func (e CustomError) Unwrap() error {
	return e.Wrapped
}

func SomeFunction() error {
	return CustomError{
		Message: "original error: something went wrong",
		Wrapped: errors.New("wrapped error"),
	}
}
