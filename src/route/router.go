package route

import (
	"conversion-currency/src/route/routers"

	"github.com/gorilla/mux"
)

func RunRoutesInfo() *mux.Router {
	r := mux.NewRouter()

	return routers.ConfigRoute(r)
}
