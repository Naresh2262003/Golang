package main

import (
	"fmt"
	"sync"
)

func worker(a int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("worker", a, "started")
	// some task done
	fmt.Println("worker", a, "ended")
}

func main() {
	fmt.Println("Exploring go routines started")

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// wait for all the go routines
	wg.Wait()
	fmt.Println("workers task completed")
}
