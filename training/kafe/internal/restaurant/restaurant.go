package restaurant


type FoodItem interface {
	GetPrice() float64
}

type Restaurant struct{
	Menu []FoodItem
}

func (r *Restaurant)AddToMenu(item FoodItem){
	r.Menu = append(r.Menu, item)

}
func (r *Restaurant)GetMenuPrice()float64{
	var sum float64
	for _,v:=range r.Menu{
		sum+=float64(v.GetPrice())
	}
	return sum
}