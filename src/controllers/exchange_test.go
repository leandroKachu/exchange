package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestConverter(t *testing.T) {
	// Crie um request falso com os parâmetros necessários para a função Converter
	req, err := http.NewRequest("GET", "/exchange/10/BRL/USD/4.50", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	Converter(rr, req)

	// Verifique o código de status da resposta HTTP
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status code esperado %v, mas recebido %v", http.StatusOK, status)
	}

	// Verifique o conteúdo da resposta JSON
	expected := `{"valorConvertido":45,"simboloMoeda":"$"}`
	if rr.Body.String() != expected {
		t.Errorf("resposta esperada %v, mas recebida %v", expected, rr.Body.String())
	}
}
