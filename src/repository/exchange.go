package repository

import (
	"database/sql"
)

type Conversion struct {
	db *sql.DB
}

func NewRepositoryConversion(db *sql.DB) *Conversion {
	return &Conversion{db}
}

type ConversionResult struct {
	ConvertedValue float64 `json:"valorConvertido"`
	CurrencySymbol string  `json:"simboloMoeda"`
}

func (repository Conversion) CreateConversion(ExchangeRate, ConvertedValue float64, FromCurrency, CurrencySymbol, ToCurrency string) (interface{}, error) {

	row := repository.db.QueryRow("SELECT DISTINCT id, converted_value, currency_symbol FROM conversions WHERE from_currency = ? AND to_currency = ? AND exchange_rate = ?", FromCurrency, ToCurrency, ExchangeRate)

	var convertedValue float64
	var currencySymbol string
	var id int64

	err := row.Scan(&id, &convertedValue, &currencySymbol)
	if err == nil {
		result := ConversionResult{
			ConvertedValue: convertedValue,
			CurrencySymbol: currencySymbol,
		}

		return result, nil
	} else if err != sql.ErrNoRows {
		return 0, err
	}

	statement, erro := repository.db.Prepare(
		"INSERT INTO conversions (from_currency, to_currency, exchange_rate, converted_value, currency_symbol) VALUES ( ?, ?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	_, erro = statement.Exec(FromCurrency, ToCurrency, ExchangeRate, ConvertedValue, CurrencySymbol)
	if erro != nil {
		return 0, erro
	}

	result := ConversionResult{
		ConvertedValue: ConvertedValue,
		CurrencySymbol: CurrencySymbol,
	}

	return result, nil

}
