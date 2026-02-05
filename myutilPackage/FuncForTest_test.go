package myutil

import "testing"

func TestFuncForTest(t *testing.T) {
	result, _ := FuncForTest(3, 6, '%')
	if result != 3 {
		t.Errorf("Result was incorrect: Got %d, want %d", result, 9)
	}
}
