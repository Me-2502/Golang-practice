package main

import (
	"fmt"
)

func main() {
	// BASICS
	// fmt.Println("Hello world!")
	// var temp int = 9
	// fmt.Scan(&temp)
	// fmt.Printf("%d, %T\n", temp, temp)
	// reader := bufio.NewReader(os.Stdin)
	// name, _ := reader.ReadString('\n')
	// fmt.Println(name)
	// fmt.Scan(&name)
	// fmt.Println(name)

	// ARRAYS
	// numbers := [6]int{3, 4, -2, 92, 46}
	// fmt.Println(numbers[3], numbers[5])
	// var fruits [5]string
	// fruits[0] = "Apple"
	// fruits[4] = "Banana"
	// fmt.Println(fruits, len(fruits))

	// PACKAGES
	// PrintHello()
	// myutil.Conditionals()
	// myutil.Functions()
	// myutil.Slices()
	// myutil.Structs()
	// myutil.Interfaces()
	// myutil.Maps()

	// s := []int{1, 2, 3, 4, 5}
	// a := []int{1, 2, 3, 4, 5}
	// fmt.Printf("%T %T\n", s, a)
	// nums := make([]int, 3, 6)
	// fmt.Println(nums, len(nums), cap(nums))
	// nums = append(nums, 2, 92, 02, -24)
	// fmt.Println(nums, len(nums), cap(nums))

	var point *int
	fmt.Println(point)
	a := 34
	point = &a
	var pop **int = &point
	fmt.Println(pop, *pop, **pop)
	var popop ***int = &pop
	fmt.Println(popop, *popop, **popop)
}
