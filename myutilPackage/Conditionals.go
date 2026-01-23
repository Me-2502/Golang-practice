package myutil

import "fmt"

func Conditionals() {
	// IF - ELSE-IF - ELSE
	fmt.Println("IF - ELSE-IF - ELSE")
	age := 19
	if age <= 14 {
		fmt.Println("You are a kid.")
	} else if age <= 18 {
		fmt.Println("You are a kid.")
	} else if age <= 60 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}
	fmt.Println()

	// SWITCH
	fmt.Println("SWITCH")
	x := 9
	fmt.Println("The number x is", x)
	switch x {
	case 1, 2, 3, 4, 5:
		fmt.Println("The number is smaller than equal to 5.")
	case 6, 7, 8, 9, 10:
		fmt.Println("The number is smaller than equal to 10.")
	default:
		fmt.Println("The number is greater than 10.")
	}
	fmt.Println()

	stage := 0
	fmt.Println("Initial stage is", stage)
	switch stage {
	case 0:
		fmt.Println("You have entered the first stage")
		// stage++
		fallthrough
	case 1:
		fmt.Println("You have entered the second stage")
		stage++
		if stage == 0 {
			break
		}
		fallthrough
		// break
	case 2:
		fmt.Println("You have entered the third stage")
		stage++
		// fallthrough
	case 3:
		fmt.Println("You have entered the final stage")
	}
	fmt.Println()
}
