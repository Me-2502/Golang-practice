package myutil

import "errors"

func FuncForTest(a, b int, o byte) (ans int, err error) {
	switch o {
	case '+':
		ans = a + b
	case '-':
		ans = a - b
	case '*':
		ans = a * b
	case '/':
		if b == 0 {
			err = errors.New("Can't divide by zero")
		} else {
			ans = a + b
		}
	default:
	}
	return ans, err
}
