package repository

import (
	"database/sql"
	"fmt"
)

type Conversion struct {
	db *sql.DB
}

func NewRepositoryConversion(db *sql.DB) *Conversion {
	return &Conversion{db}
}

type ConversionResult struct {
	ValorConvertido float64 `json:"valorConvertido"`
	SimboloMoeda    string  `json:"simboloMoeda"`
}

func (repository Conversion) CreateConversion(amount, ExchangeRate, ConvertedValue float64, FromCurrency, CurrencySymbol, ToCurrency string) (interface{}, error) {

	row := repository.db.QueryRow("SELECT DISTINCT id, converted_value, currency_symbol FROM conversions WHERE amount = ? AND from_currency = ? AND to_currency = ? AND exchange_rate = ?", amount, FromCurrency, ToCurrency, ExchangeRate)

	var convertedValue float64
	var currencySymbol string
	var id int64

	err := row.Scan(&id, &convertedValue, &currencySymbol)
	fmt.Println(err)

	if err == nil {
		result := ConversionResult{
			ValorConvertido: convertedValue,
			SimboloMoeda:    currencySymbol,
		}

		return result, nil
	} else if err != sql.ErrNoRows {
		// Ocorreu um erro ao executar a consulta
		return 0, err
	}

	statement, erro := repository.db.Prepare(
		"INSERT INTO conversions (amount, from_currency, to_currency, exchange_rate, converted_value, currency_symbol) VALUES (?, ?, ?, ?, ?, ?)",
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	_, erro = statement.Exec(amount, FromCurrency, ToCurrency, ExchangeRate, ConvertedValue, CurrencySymbol)
	if erro != nil {
		return 0, erro
	}

	result := ConversionResult{
		ValorConvertido: ConvertedValue,
		SimboloMoeda:    CurrencySymbol,
	}

	return result, nil

}
