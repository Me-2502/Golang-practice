package myutil

import (
	"errors"
	"fmt"
)

func calculator(a, b int, opr string) (int, error) {
	switch opr {
	case "ADD":
		return a + b, nil
	case "MIN":
		return a - b, nil
	case "MUL":
		return a * b, nil
	case "DIV":
		if b == 0 {
			return 0, errors.New("Can't divide by zero!")
		}
		return a / b, nil
	}
	return 0, errors.New("Can't find valid operation!")
}

func Functions() {
	// FUNCTIONS
	fmt.Println("Functions")
	fmt.Println()

	// Calculator
	var temp int
	var err error
	temp, err = calculator(5, 4, "ADD")
	if err == nil {
		fmt.Println("5 + 4 =", temp)
	} else {
		fmt.Println(err)
	}
	temp, err = calculator(5, 4, "MOD")
	if err == nil {
		fmt.Println("5 % 4 =", temp)
	} else {
		fmt.Println(err)
	}
	fmt.Println()
}
