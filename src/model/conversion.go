package model

type Conversion struct {
	Amount         float64 `json:"amount"`
	FromCurrency   string  `json:"from"`
	ToCurrency     string  `json:"to"`
	ExchangeRate   float64 `json:"rate"`
	ConvertedValue float64 `json:"valorConvertido"`
	CurrencySymbol string  `json:"simboloMoeda"`
}
