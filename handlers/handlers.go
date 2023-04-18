package handlers

import (
	"net/http"

	mw "github.com/jnnkrdb/gomw/middlewares"
)

type HttpFunctionSet []HttpFunction

type HttpFunction struct {
	Pattern     string
	Function    func(http.ResponseWriter, *http.Request)
	Middlewares mw.MiddleWareChain
}

// return the aggregated http.Handler
func GetHandler(hfs []HttpFunction) http.Handler {
	var h = http.NewServeMux()
	for i := range hfs {
		h.Handle(hfs[i].Pattern, hfs[i].Middlewares.ThenFunc(hfs[i].Function))
	}
	return h
}
