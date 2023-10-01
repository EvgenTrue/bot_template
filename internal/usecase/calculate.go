package usecase

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"strconv"

  )

type Response struct{
	Status int  `json:status`
	Message string `json:message`
	Data map[string]string `json:data`
}


type CalculateCurrencyUsecase struct{
	
}
func New()*CalculateCurrencyUsecase{
	return &CalculateCurrencyUsecase{

	}
}
func (c *CalculateCurrencyUsecase)Calculate(sum int, currency string)(int, error){
	res, err := http.Get(fmt.Sprintf("https://currate.ru/api/?get=rates&pairs=%sRUB&key=c928d94ff04728cf9ff8ac5e77ec17de", currency))
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body), err)


	var s Response  
	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println(err)
	    return 0, err
	}
	fmt.Println(s)
	k:=fmt.Sprintf("%sRUB",currency)
	fmt.Println(s.Data[k])
	 
	
	t, err := strconv.ParseFloat(s.Data[k], 64);
	 if err != nil {
    	fmt.Println(err) 
		return 0,err
	
	}

	t=float64(sum) *t*1000
	i := int(t)   
		fmt.Println(i)
	return i, nil 
}

//{"status":200,"message":"rates","data":{"USDRUB":"64.1824"}}