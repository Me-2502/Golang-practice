package myutil

import (
	"errors"
	"fmt"
	"unsafe"
)

type transformType func(int) int

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

func doubleNumbers(nums *[]int) {
	fmt.Println(unsafe.Sizeof(nums))
	for i, v := range *nums {
		(*nums)[i] = v << 1
	}
}

func duplicateDoubleNumbers(nums []int) {
	fmt.Println(unsafe.Sizeof(nums))
	for i, v := range nums {
		nums[i] = v << 1
	}
}

func transformNumbers(numbers []int, transform transformType) {
	for i, v := range numbers {
		numbers[i] = transform(v)
	}
}

func double(i int) int {
	return i * 2
}

func triple(i int) int {
	return i * 3
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func variaticSum(message string, nums ...int) (ans int) {
	fmt.Println(message)
	nums[0] = -34
	for _, v := range nums {
		ans += v
	}
	return
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

	numbers := []int{1, 3, 4, 6, -12}
	doubleNumbers(&numbers)
	fmt.Println(numbers)
	duplicateDoubleNumbers(numbers)
	fmt.Println(numbers)
	transformNumbers(numbers, triple)
	fmt.Println(numbers)
	transformNumbers(numbers, func(i int) int { return i / 6 })
	fmt.Println(numbers)
	transformNumbers(numbers, createTransformer(6))
	fmt.Println(numbers)
	fmt.Println()

	fmt.Println(variaticSum("Added using variatic function", 3, 53, -35, 654))
	fmt.Println(variaticSum("Again added using variatic function", numbers...))
	fmt.Println(numbers)
	fmt.Println()

	// var j int
	for j := 0; j < 3; j++ {
		defer func(j int) {
			fmt.Println(j)
		}(j)
	}
}
