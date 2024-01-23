package provider

type DammyProvider struct {
}

type ResponseDammy struct {
	Success bool              `json:"success"`
	Rates   map[string]string `json:"rates"`
}

func NewDammyProvider() *DammyProvider {
	return &DammyProvider{}
}
func (c *DammyProvider) GetName() string {  //???
	return "fixer"
}
func (c *DammyProvider) GetCurrency(currency string) (float64, error) {
	return 60.24, nil
}
