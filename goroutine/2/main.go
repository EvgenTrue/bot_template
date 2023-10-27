// 2. Горутины для вывода сообщений
// messages := []string{"Привет", "Здравствуй", "Приветствую"}
// Есть слайс, надо вывести каждое сообщение в отдельной горутине через n секунд, где n это индекс сообщения в слайсе
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	messages := []string{"Привет", "Здравствуй", "Приветствую", "ewrwwer", "wdwed", "weqwwwee"}
	var wg sync.WaitGroup
	wg.Add(len(messages))
	var mu sync.RWMutex
	for i, v := range messages {

		go func(i int, v string) {
			mu.RLock()
			defer mu.RUnlock()
			time.Sleep(time.Second * time.Duration(i))
			fmt.Println(v)
			wg.Done()
		}(i, v)

	}
	wg.Wait()
}
