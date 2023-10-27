//* Параллельный подсчет суммы: Напишите программу, которая использует две горутины для параллельного подсчета суммы элементов в двух половинах среза целых чисел.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var hum int
	var sum int
	slice:=[]int{1,2,4,1,1,1,1,2}
	ch:=make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	var mu sync.RWMutex

	go func(){
		mu.Lock()
		defer mu.Unlock()
		for _,v:=range slice[0:4]{
			sum+=v
			ch<-v
			time.Sleep(time.Second *1)
			fmt.Println("первая половина:",sum)
		}
		wg.Done()
	}()
	
	go func(){
		mu.Lock()
		defer mu.Unlock()
		for _,v:=range slice[4:8]{
			hum+=v
			ch<-v
			time.Sleep(time.Second *1)
			fmt.Println("вторая половина:",hum)
		}
		wg.Done()
	}()
	time.Sleep(time.Second *5)

	go func(){
		wg.Wait()
		close(ch)

	}()
	for v:=range ch{
		fmt.Println(v)
	}
	fmt.Println(slice)
}
