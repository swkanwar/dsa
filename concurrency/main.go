package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var counter int

	var mutex sync.Mutex

	numGoroutines := 10000
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
			fmt.Println(counter)
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
