package provider

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type FixerProvider struct {
}

type ResponseFixer struct {
	Success bool              `json:"success"`
	Rates   map[string]string `json:"rates"`
}

func NewFixerProvider() *FixerProvider {
	return &FixerProvider{}
}

func (c *FixerProvider) GetCurrency(currency string) (float64, error) {
	res, err := http.Get(fmt.Sprintf("https://data.fixer.io/api/latest?&base=RUB&symbols=%s&access_key=bff0d81f6bc0c7e5c279c24bb5f2bd0e", currency))
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body), err)

	var s ResponseFixer
	err = json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	fmt.Println(s)
	k := fmt.Sprintf("%sRUB", currency)
	fmt.Println(s.Rates[k])

	t, err := strconv.ParseFloat(s.Rates[k], 64)
	if err != nil {
		fmt.Println(err)
		return 0, err

	}
	return t, nil
}
