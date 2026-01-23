package myutil

import "fmt"

func addElement(s []int, elem int) {
	// Bugged
	s = append(s, elem)
	fmt.Println("Inside function:", s, len(s), cap(s))
	// Fixed
	// return append(s, elem)
}

func addElementViaRef(s *[]int, elem int) {
	*s = append(*s, elem)
}

func abnormalBehaviour() {
	slice := []int{1, 2, 3}
	// Bugged
	addElement(slice, 4)
	// Fixed
	// slice = addElement(slice, 4)
	fmt.Println("Outside function:", slice, len(slice), cap(slice))
	// fmt.Println("Outside function:", slice[3])

	slice = make([]int, 3, 5)
	slice[0], slice[1], slice[2] = 1, 2, 3
	addElement(slice, 4)
	fmt.Println("Outside function:", slice, len(slice))
	newSlice := slice[:4]
	fmt.Println("New slice:", newSlice, len(newSlice))

	// Edge case
	original := []int{1, 2, 3, 4, 5}
	s1 := original[:3]
	s2 := original[:3]
	s1 = append(s1, 999)
	// Fix
	// s1 := original[:3:3]
	// s1 = append(s1, 999)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(original)

	// Nill
	var nilSlice []int             // nil slice
	emptySlice := []int{}          // empty slice
	madeSlice := make([]int, 0)    // empty slice
	fmt.Println(nilSlice == nil)   // true
	fmt.Println(emptySlice == nil) // false
	fmt.Println(madeSlice == nil)  // false
}

func badFunc() []byte {
	buf := make([]byte, 1024)
	return buf[:10]
}

func Slices() {
	// SLICES
	s := []int{3, 5, 6}
	fmt.Println(len(s), cap(s))
	s = s[:2]
	fmt.Println(s, len(s), cap(s))
	s = s[:3]
	fmt.Println(s, len(s), cap(s))
	var c = make([]int, 4, 5)
	copy(c, s)
	fmt.Println(c, len(c), cap(c))
	c[1] = 3223
	fmt.Println(c, s)
	x := []int{1, 2, 3, 4, 5}
	y := x[1:3]
	z := x[1:3:3]
	fmt.Println(cap(y), cap(z))
	x = append(x[:2], x[3:]...)
	fmt.Println(x, len(x), cap(x))
	spaceBlocker := badFunc()
	fmt.Println(len(spaceBlocker), cap(spaceBlocker))
	fmt.Println()

	abnormalBehaviour()

	// Good practices
	// type BigStruct struct {
	// 	data [1000]int
	// }
	// bigSlice := make([]BigStruct, 1000)
	// // Slow - copies each 8KB struct
	// for _, v := range bigSlice {
	// 	process(v)  // v is a copy!
	// }
	// // Fast - no copies
	// for i := range bigSlice {
	// 	process(&bigSlice[i])  // pointer, no copy
	// }

	fmt.Println("Hello world from myutil.go")
}
