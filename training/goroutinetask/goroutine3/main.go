// * Параллельный поиск минимума и максимума: Создайте программу, которая использует две горутины для параллельного поиска минимального и максимального значения в срезе чисел.
package main

import (
	"fmt"
	"sort"
	"time"
)

func main(){
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	
	go func (slise[]int) {
		sort.Ints(slise)
		fmt.Println(slise[0])
		
		
	}(slice)
	go func (slise[]int) {
		sort.Ints(slise)
		fmt.Println(slise[len(slise)-1])
		
	}(slice)
	time.Sleep(time.Second * 5)
}