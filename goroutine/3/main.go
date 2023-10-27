//3*.Параллельная обработка слайса чисел версия 2:
//Условие такое же как в задаче 1, только слайс надо заменить на функцию генератор, которая возвращает
//канал и в канал закидывает числа.
//ch := generator(1000) <- кладет числа от 0 до 1000 в канал

package main

import (
	"fmt"
	"sync"
)

func generator(number int) <-chan int {
	ch := make(chan int)

	go func() {
		for i := 0; i <= number; i++ {
			ch <- i
		}

		close(ch)

	}()

	return ch

}

func main() {
	ch := generator(5)
	var wg sync.WaitGroup
	wg.Add(3)
	var mu sync.RWMutex

	go func() {
		mu.RLock()
		defer mu.RUnlock()
		var sum int
		for v := range ch {
			sum += v
		}
		fmt.Println("сумма:", sum)
		wg.Done()
	}()

	go func() {
		mu.RLock()
		defer mu.RUnlock()
		var min int
		for v := range ch {
			if min == 0 {
				min = v
				continue
			}
			if min > v {
				min = v
			}
		}
		fmt.Println("минимум", min)
		wg.Done()
	}()

	go func() {
		mu.RLock()
		defer mu.RUnlock()
		var max int
		for v := range ch {
			if max < v {
				max = v
			}
		}
		fmt.Println("макс", max)
		wg.Done()
	}()
	wg.Wait()
}
