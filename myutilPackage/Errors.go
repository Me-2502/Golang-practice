package myutil

import (
	"errors"
	"fmt"
)

var invalidCredentialsErr = errors.New("Invalid credentials")
var usernameNotFoundErr = errors.New("Username was not found")
var passwordMismatchErr = errors.New("Wrong password")

var credentials = map[string]string{
	"John Doe": "#akr40sj4",
}

func login(user, pass string) error {
	originalPass, ok := credentials[user]
	if !ok {
		return errors.Join(invalidCredentialsErr, fmt.Errorf("%w: %s", usernameNotFoundErr, user))
	}
	if originalPass != pass {
		return errors.Join(invalidCredentialsErr, fmt.Errorf("%w", passwordMismatchErr))
	}
	return nil
}

func printLoginStatus(e error) {
	if e != nil {
		if errors.Is(e, passwordMismatchErr) {
			e = e.(interface {
				Unwrap() []error
			}).Unwrap()[0]
			fmt.Println(e)
		} else {
			fmt.Println(e)
		}
		return
	}
	fmt.Println("Logged in")
}

func Errors() {
	// ERRORS
	err := errors.New("Sample error!")
	fmt.Println(err)
	var e error
	e = login("Mark Smith", "#ie5ral42")
	printLoginStatus(e)
	e = login("John Doe", "#ie5ral42")
	printLoginStatus(e)
	e = login("John Doe", "#akr40sj4")
	printLoginStatus(e)
	fmt.Println()

	var s []int
	var m map[int]int
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("Recover1:", p)
		}
	}()
	defer func() {
		fmt.Println("I am between recoveries")
	}()
	defer func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("Nested recover:", p)
			}
		}()
		if p := recover(); p != nil {
			panic("ABC")
			fmt.Println("Recover2:", p)
		}
	}()

	s = append(s, 23)
	fmt.Println("I am panic free! Yay!")
	s = s[0:len(s):0]
	fmt.Println("I am panic free 2! Yay!")
	m[23] = 23

	fmt.Println("I am panic free 3! Yay!")
	fmt.Println()
}
