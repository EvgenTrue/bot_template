package storage

import "sync"

//Создать InMemoryStorage (map + lock) // Продукты 1000 RUB map + lock (* interface)

func NewMemoryStorage()*InMemoryStorage{
return &InMemoryStorage{
	List: make(map[string]Item),
}
}

type InMemoryStorage struct {
	mu sync.RWMutex
	List map[string]Item
  }

type Item struct{
	Name string
	Sum int
	Currency string
} 

func (i *InMemoryStorage)AddItem(item Item){
	i.mu.Lock()
	defer i.mu.Unlock()
	 i.List[item.Name]=item
	   
}  
func (i* InMemoryStorage)GetItem(key string) Item {
	i.mu.RLock()
	defer i.mu.RUnlock()
	return i.List[key]
}
 
