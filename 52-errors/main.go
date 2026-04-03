package main

import "fmt"

func main() {
	result, err := Divide(1, 0)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println(result)
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		//return 0, fmt.Errorf("cannot divide by zero: b = %f", b)
		return 0, CustomError{Code: 1, Message: "cannot divide by zero"}
	}
	return a / b, nil
}

type CustomError struct {
	Code    int
	Message string
}

// error interface just have an Error() string method
func (c CustomError) Error() string {
	return fmt.Sprintf("error code %d: %s", c.Code, c.Message)
}
