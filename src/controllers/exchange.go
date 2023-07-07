package controllers

import (
	"conversion-currency/src/database"
	errorsresponse "conversion-currency/src/errorsResponse"
	"conversion-currency/src/repository"
	"conversion-currency/src/services"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func Converter(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fromCurrency := strings.ToUpper(params["from"])
	toCurrency := strings.ToUpper(params["to"])
	exchangeRate, _ := strconv.ParseFloat(params["rate"], 64)

	if toCurrency == "0" || fromCurrency == "0" || exchangeRate == 0 {
		errorsresponse.Error(w, http.StatusBadRequest, errors.New("need valid parameters"))
		return
	}

	float, err := services.GetTaxCoin(fromCurrency, toCurrency)

	if err != nil {
		errorsresponse.Error(w, http.StatusInternalServerError, err)
		return
	}

	var convertedValue = exchangeRate * float
	currencySymbol := "$"

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
	datatoDB, err := repository.CreateConversion(exchangeRate, convertedValue, fromCurrency, currencySymbol, toCurrency)

	if err != nil {
		errorsresponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	returndb, err := json.Marshal(datatoDB)

	if err != nil {
		errorsresponse.Error(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(returndb)
}
