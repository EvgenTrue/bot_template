//* Параллельная проверка на простоту: Создайте программу, которая использует несколько горутин для параллельной проверки простоты нескольких чисел.

package main

import (
	"fmt"
	"time"
)

func main(){

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

go func(){
	for _,v:=range slice{
	for i := 1; i <= 10; i++ {

		if i%i == 0 && i%1 == 0  {
	
		} else {
	
		}
	}
	time.Sleep(time.Second * 1)
	fmt.Println("one:", v)
	}
		}()


go func(){
	for _,v:=range slice{
	for i := 1; i <= 10; i++ {

		if i%i == 0 && i%1 == 0 {
	
		} else {
	
		}
	}
	time.Sleep(time.Second * 1)
	fmt.Println("one:", v)
}
		}()
		time.Sleep(time.Second * 5)
}
