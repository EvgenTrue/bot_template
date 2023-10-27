//1. Создайте интерфейс FoodItem с методом GetPrice() float64, который возвращает цену пункта меню.

//2. Реализуйте структуры Dish и Beverage, представляющие блюдо и напиток соответственно. Обе структуры должны реализовать метод GetPrice().

//3. Создайте структуру Restaurant с полем menu, которое представляет собой список FoodItem. Метод AddToMenu(item FoodItem) должен добавлять пункт меню в список ресторана.

//4. Реализуйте метод GetMenuPrice() в структуре Restaurant, который вычисляет и возвращает общую стоимость всего меню ресторана.

package main

import ("fmt"

"github.com/EvgenTrue/kafe/internal/restaurant"
"github.com/EvgenTrue/kafe/internal/menu"
)
 

 

func main(){
	itemchicken:=&menu.Dish{Name: "chiken", Price: 12.32}
	restaurant:=restaurant.Restaurant{}
	restaurant.AddToMenu(itemchicken)
	itemkola:=&menu.Beverage{Name: "Coca-cola", Price: 5.32}
    restaurant.AddToMenu(itemkola)
	itogo:=restaurant.GetMenuPrice()
	fmt.Println(itogo)
	
}


