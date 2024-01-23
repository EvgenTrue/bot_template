package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type CurrateProvider struct {				// создаем структуру провайдера
}

type ResponseCurrate struct {					// с сайта провайдера привязываемся по полям 
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Data    map[string]string `json:"data"`
}

func NewCurrateProvider() *CurrateProvider {   // создаем функцию для обращения к провайдеру и обратной связи
	return &CurrateProvider{}
}
func (c *CurrateProvider)GetName()string{    // функция которая создает 
	return "currency"
} 
func (c *CurrateProvider) GetCurrency(currency string) (float64, error) {  //функция для запроса курса 
	res, err := http.Get(fmt.Sprintf("https://currate.ru/api/?get=rates&pairs=%sRUB&key=c928d94ff04728cf9ff8ac5e77ec17de", currency))
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body), err)

	var s ResponseCurrate
	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(s)
	k := fmt.Sprintf("%sRUB", currency)
	fmt.Println(s.Data[k])

	t, err := strconv.ParseFloat(s.Data[k], 64)
	if err != nil {
		fmt.Println(err)
		return 0, err

	}
	return t, nil
}
