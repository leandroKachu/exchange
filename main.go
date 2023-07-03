package main

import (
	"conversion-currency/src/config"
	"conversion-currency/src/route"
	"fmt"
	"log"
	"net/http"
)

// Main contem configuracoes basicas para execusao do sistema raiz.
func main() {
	config.InitBaseInfos()
	r := route.RunRoutesInfo()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
