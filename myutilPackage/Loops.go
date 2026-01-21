package myutil

import "fmt"

func Loops() {
	// LOOPS
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	temp := 0
	for {
		temp += 2
		if temp > 10 {
			break
		}
		fmt.Println(temp)
	}

	arr := [5]int{100, 200, 300, 400}
	for i, v := range arr {
		fmt.Printf("Value at index %d is %d\n", i, v)
	}
	ages := map[string]int{"Alice": 25, "Bob": 30}
	for key, value := range ages {
		fmt.Printf("%s is %d years old\n", key, value)
	}
	fmt.Println()

	// ANOMALUS BEHAVIOUR
	numbers := []int{1, 2, 3}
	for _, v := range numbers {
		numbers = append(numbers, 999)
		fmt.Println(v)
	}
	text := "Hello, 世界"
	for i, char := range text {
		fmt.Printf("Index: %d, Char: %c\n", i, char)
	}
	fmt.Println()
}
