package controllers

import (
	"conversion-currency/src/database"
	errorsresponse "conversion-currency/src/errorsResponse"
	"conversion-currency/src/repository"
	"encoding/json"
	"errors"
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
	currencySymbol := "$"

	convertedValue := amount * exchangeRate

	if amount == 0 || toCurrency == "0" || fromCurrency == "0" || exchangeRate == 0 {
		errorsresponse.Error(w, http.StatusBadRequest, errors.New("need valid parameters"))
		return
	}

	if toCurrency == "EUR" {
		currencySymbol = "€"
	} else if toCurrency == "BTC" {
		currencySymbol = "฿"
	}

	db, err := database.Connection()
	if err != nil {
		errorsresponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repository.NewRepositoryConversion(db)
	result, err := repository.CreateConversion(amount, exchangeRate, convertedValue, fromCurrency, currencySymbol, toCurrency)
	if err != nil {
		errorsresponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	jsonData, err := json.Marshal(result)
	if err != nil {
		errorsresponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
