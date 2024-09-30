package handlers

import (
	"net/http"

	mw "github.com/jnnkrdb/go-helpmepls/httpmw/middlewares"
)

type HttpFunctionSet []HttpFunction

type HttpFunction struct {
	Pattern     string
	MainHandler http.Handler
	Middlewares mw.MiddleWareChain
}

// return the aggregated http.Handler
func GetHandler(hfs []HttpFunction) http.Handler {
	var h = http.NewServeMux()
	for i := range hfs {
		h.Handle(hfs[i].Pattern, hfs[i].Middlewares.Then(hfs[i].MainHandler))
	}
	return h
}
