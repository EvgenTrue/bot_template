//* Параллельная генерация простых чисел: Напишите программу, которая использует несколько горутин для параллельной генерации простых чисел в разных диапазонах.

package main

import (
	"fmt"
	"time"
)

func main(){

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

go func(){
	for _,i:=range slice{
	 

		if i%i == 0 && i%1 == 0 && i>=5 && i<=10{
			fmt.Println("one:",i)}
	 
	}
}()
		 
	
		 
	
		 


go func(){
	for _,v:=range slice{
	for i := 1; i <= 10; i++ {

		if i%i == 0 && i%1 == 0 && i>=0 && i<5{
	
		} else {
	
		}
	}
	time.Sleep(time.Second * 1)
	fmt.Println("two:", v)
}
		}()
		time.Sleep(time.Second * 5)
}
}