package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"unsafe"
)

var i int = 1

func print(s string, doneChan chan int, errChan chan error, c bool) {
	fmt.Println(s)
	if !c {
		errChan <- errors.New("Sample error!")
	}
	if c {
		i++
		doneChan <- (i)
		// close(doneChan)
	}
}

func Concurrency() {
	// CONCURRENCY
	done := make(chan int, 4)
	errorChan := make(chan error)
	fmt.Println(unsafe.Sizeof(make(chan int, 5)))
	go print("Hi from concurrency!", done, errorChan, false)
	go print("Hello world from concurrency!", done, errorChan, true)
	go print("Hello world from concurrency!", done, errorChan, true)
	go print("Hello world from concurrency!", done, errorChan, true)
	go print("Hello world from concurrency!", done, errorChan, true)
	for j := 0; j < 5; j++ {
		select {
		case err := <-errorChan:
			fmt.Println(err)
		case i := <-done:
			fmt.Println(i)
		case <-time.After(time.Microsecond * 10):
			fmt.Println("Time limit exceeded")
		}
	}
	fmt.Println()

	nums := []int{1, -2, 492, 0, 24}
	sliceInputChannel := sliceToChannel(nums)
	squaredInputChannel := squareFunc(sliceInputChannel)
	for v := range squaredInputChannel {
		fmt.Println(v)
	}
	fmt.Println()

	numbers := []int{2, 9394, 35, -687, 0}
	for len(numbers) > 1 {
		sliceElementToChannel := sliceToChannel(numbers)
		halfLen := (len(numbers) - 1) / 2
		for i := 0; i <= halfLen; i++ {
			addChannel := addFunc(sliceElementToChannel)
			numbers[i] = <-addChannel
			// fmt.Println(numbers[i])
		}
		// fmt.Println(numbers)
		numbers = numbers[0 : (len(numbers)+1)/2]
	}
	fmt.Println(numbers[0])
	fmt.Println()

	var wg sync.WaitGroup
	fmt.Println(unsafe.Sizeof(wg))
	wg.Add(1)
	fmt.Println("Before done")
	go func() {
		fmt.Println("Done")
		wg.Done()
	}()
	go func() {
		fmt.Println("Extra")
	}()
	fmt.Println("After done")
	// for i := 0; i < 10; i++ {
	// 	i++
	// }
	wg.Wait()
	fmt.Println()

	p := sync.Pool{
		New: func() any {
			// return new(*bytes.Buffer)
			return nil
		},
	}
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println()

	var c sync.Cond
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println()
}

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			out <- v
		}
	}()
	return out
}

func addFunc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		a, ok := <-in
		if !ok {
			return
		}
		b, ok := <-in
		if !ok {
			out <- a
			return
		}
		out <- a + b
	}()
	return out
}

func squareFunc(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for v := range in {
			out <- v * v
		}
		close(out)
	}()
	return out
}
