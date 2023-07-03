package controllers

import (
	"conversion-currency/src/database"
	"conversion-currency/src/repository"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Converter(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	amount, _ := strconv.ParseFloat(params["amount"], 64)
	fromCurrency := params["from"]
	toCurrency := params["to"]

	exchangeRate, _ := strconv.ParseFloat(params["rate"], 64)

	convertedValue := amount * exchangeRate
	currencySymbol := "$"

	if toCurrency == "EUR" {
		currencySymbol = "€"
	} else if toCurrency == "BTC" {
		currencySymbol = "฿"
	}

	db, erro := database.Connection()
	if erro != nil {
		// respostas.Erro(w, http.StatusInternalServerError, erro)
		fmt.Println("deu error: ", erro)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryConversion(db)
	result, err := repository.CreateConversion(amount, exchangeRate, convertedValue, fromCurrency, currencySymbol, toCurrency)
	if err != nil {
		fmt.Println(err)
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
