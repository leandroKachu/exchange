package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI    string
	Method string
	Func   func(http.ResponseWriter, *http.Request)
}

func ConfigRoute(r *mux.Router) *mux.Router {
	routerExchange := routerExchange
	r.HandleFunc(routerExchange.URI, routerExchange.Func).Methods(routerExchange.Method)
	return r
}
