package routers

import (
	"conversion-currency/src/controllers"
	"net/http"
)

var routerExchange = Route{
	URI:    "/exchange/{amount}/{from}/{to}/{rate}",
	Method: http.MethodPost,
	Func:   controllers.Converter,
}
