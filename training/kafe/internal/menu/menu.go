package menu

type Dish struct{
	Name string
	Price float64 
}
type Beverage struct{
	Name string
	Price float64
}

func (d *Dish)GetPrice()float64{
return d.Price
}

func (b *Beverage)GetPrice()float64{
return b.Price
}
