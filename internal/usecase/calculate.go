package usecase

import (
	"fmt"
)

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}
type CurrencyProvider interface {
	GetCurrency(currency string) (float64, error)
	GetName()string
}

type CalculateCurrencyUsecase struct {
	 
	provider []CurrencyProvider
}

func New(p []CurrencyProvider) *CalculateCurrencyUsecase {
	return &CalculateCurrencyUsecase{
		provider: p,
	}
}
func (c *CalculateCurrencyUsecase) Calculate(sum int, currency string) (int, error) {
	ch:=make(chan float64)
	for _,pp:=range c.provider{
		go func(){
			var t float64
			var err error
			t, err = pp.GetCurrency(currency)
	
			if err != nil {
				fmt.Println(err)
				return
			} 
			ch<-t
		}()
	
	}


	t = float64(sum) * t * 1000
	i := int(t)
	fmt.Println(i)
	return i, nil
}

//{"status":200,"message":"rates","data":{"USDRUB":"64.1824"}}

// разбить калькулятор на провайдеры, и разбить на калькулятор
// написать новый проваййдер.
// иметь возможность выбрать в мейне провайдера
// тесты к провайдерам и калькулятору
