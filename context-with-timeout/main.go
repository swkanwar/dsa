package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	numGoRoutines := 5

	for i := 1; i <= numGoRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("goroutine done")
					return
				default:
					// simulate work
					fmt.Println("goroutine", i)
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("All goroutines:", numGoRoutines)

}
