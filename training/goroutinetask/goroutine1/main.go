//* Параллельный вывод четных и нечетных: Напишите программу, которая использует две горутины для параллельного вывода четных и не четеных из слайса

package main

import (
	"fmt"
	"sync"
	 
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ch:=make(chan int)
	 var wg sync.WaitGroup
	wg.Add(2)
	var mu sync.RWMutex
	
	go func() {
		mu.Lock()
		defer mu.Unlock()
		for i, v := range slice {
			if v%2 == 0 {
				 
			 ch<-v
			slice[i]=0
			}
		}
		wg.Done()
	}()

	go func() {
		mu.Lock()
		defer mu.Unlock()
		for i, v := range slice {
			if v%2 != 0 {
				  
				ch<-v
				slice[i]=0
			}
		}
		wg.Done()
	}()

	go func ()  {
		wg.Wait()
		close(ch)
	}()
	for v:= range ch{
		fmt.Println(v)
		
	}
	  fmt.Println(slice)
	 
	}
		
	 
