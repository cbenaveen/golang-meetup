package main

import (
	"errors"
	"fmt"
)

func addNumbers(no1, no2 int) int {
	return no1 + no2
}

func divide(no1, no2 int) (int, error) {
	if no2 <= 0 {
		return -1, errors.New("divide by zero")
	}

	return no1 / no2, nil
}

func main() {
	fmt.Println("Result of 10 + 20 is", addNumbers(10, 20))

	var divResult, err = divide(10, 20)
	fmt.Println("Result of 10 / 20 is", divResult, err)

	divResult, err = divide(10, 0)
	fmt.Println("Result of 10 / 0 is", divResult, err)
}
