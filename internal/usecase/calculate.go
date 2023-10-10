package usecase

import (
	"fmt"

	"github.com/evgentrue/bot_template/internal/usecase/provider"
)

type Response struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

type CalculateCurrencyUsecase struct {
	provider *provider.CurrateProvider
}

func New(p *provider.CurrateProvider) *CalculateCurrencyUsecase {
	return &CalculateCurrencyUsecase{
		provider: p,
	}
}
func (c *CalculateCurrencyUsecase) Calculate(sum int, currency string) (int, error) {
	t, err := c.provider.GetCurrency(currency)
	if err != nil {
		return 0, err
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
