package main

import (
	"fmt"
	"sync"
)

func main() {
	var max int
	var wg sync.WaitGroup

	var mu sync.RWMutex

	for i := 1000; i > 0; i-- {
		i := i
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			if i%2 == 0 && i > max {
				max = i
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("Maximum is %d", max)
}

//wr group
//lock
