//1.Параллельная обработка слайса чисел:
//Есть слайс чисел, и нужно параллельно вычислить сумму, минимум и максимум этого слайса.

package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{1, 2, 5, 23, 23, 324, 21, 64, 13}
	var wg sync.WaitGroup
	wg.Add(3)
	var mu sync.RWMutex

	go func() {
		mu.RLock()
		defer mu.RUnlock()
		var sum int
		for _, v := range slice {
			sum += v
		}
		fmt.Println(sum)
		wg.Done()
	}()

	go func() {
		mu.RLock()
		defer mu.RUnlock()
		var min int
		for _, v := range slice {
			if min == 0 {
				min = v
				continue
			}
			if min > v {
				min = v
			}
		}
		fmt.Println(min)
		wg.Done()
	}()

	go func() {
		mu.RLock()
		defer mu.RUnlock()
		var max int
		for _, v := range slice {
			if max < v {
				max = v
			}
		}
		fmt.Println(max)
		wg.Done()
	}()
	wg.Wait()
}
