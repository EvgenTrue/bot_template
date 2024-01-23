package storage

import "sync"

//Создать InMemoryStorage (map + lock) // Продукты 1000 RUB map + lock (* interface)

func NewMemoryStorage() *InMemoryStorage {    // создаем функцию которая будет хранить все в мапе 
	return &InMemoryStorage{
		List: make(map[string]Item),
	}
}

type InMemoryStorage struct {
	mu   sync.RWMutex              // создаем мютекс для избежания паники
	List map[string]Item
}

type Item struct {
	Name     string `db:"name"`
	Sum      int `db:"sum"`
	Currency string `db:"currency"`
	Id int `db:"id"`
}

func (i *InMemoryStorage) AddItem(item Item) {
	i.mu.Lock()                                              // закрываем передачу данных здля записи и чтения
	defer i.mu.Unlock()											// открываем возможность для записи и чтения
	i.List[item.Name] = item									// записываем данные в мапу

}
func (i *InMemoryStorage) GetItem(key string) Item {			// поиск по ключу из мапы 
	i.mu.RLock()                                                  // закрываем доступ для записи и чтения 
	defer i.mu.RUnlock()											// открываем доступ на чтение 
	return i.List[key]												// возвращаем значение по ключу
}
