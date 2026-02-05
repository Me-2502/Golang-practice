package main

import (
	"fmt"
	"sync"
)

func worker(wg *sync.WaitGroup, jobs <-chan int, results chan<- int) {
	defer wg.Done()
	for i := range jobs {
		// i := <-jobs
		// fmt.Println(i, "Worker")
		results <- (i * i)
	}
}

func printResults(wg *sync.WaitGroup, results <-chan int) {
	defer wg.Done()
	for i := range results {
		fmt.Println(i)
	}
}

func workerPool() {
	jobs := make(chan int)
	results := make(chan int)
	var wg, resultsWG sync.WaitGroup
	wg.Add(3)
	resultsWG.Add(1)
	for i := 0; i < 3; i++ {
		go worker(&wg, jobs, results)
	}
	go printResults(&resultsWG, results)
	for i := 1; i < 8; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	close(results)
	// wg.Add(1)
	// go func() {
	// 	for {
	// 		fmt.Println("Hello world!")
	// 	}
	// 	// wg.Done()
	// }()
	// wg.Wait()
	resultsWG.Wait()
	fmt.Println()
}

// nil channel
// Gracefull shutdown
// Range keyword
// // Functions x(i, j) vs y:=x(i,j)
// // Buffered vs Unbuffered channels
// Recursive func in defer after recover
// defer struct
// // run main package function from other files without compililng them
