package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	m := map[string]int{"a": 10, "b": 2, "c": 1}
	for _, v := range s {
		for _, j := range m {
			if v != j {
				fmt.Println(v, j)
			}
		}
	}

}
