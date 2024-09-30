package middlewares

import "net/http"

// construction of middlewares
type MiddleWare func(http.Handler) http.Handler

type MiddleWareChain struct {
	middlewares []MiddleWare
}

// create a new list of middlewares
func New(middlewares ...MiddleWare) MiddleWareChain {
	return MiddleWareChain{append(([]MiddleWare)(nil), middlewares...)}
}

// use another handler
func (mwc MiddleWareChain) Then(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}
	for i := range mwc.middlewares {
		h = mwc.middlewares[len(mwc.middlewares)-1-i](h)
	}
	return h
}

// use a func instead of a handler
func (mwc MiddleWareChain) ThenFunc(f http.HandlerFunc) http.Handler {
	if f == nil {
		return mwc.Then(nil)
	}
	return mwc.Then(f)
}

// append new middlewares to the chain
func (mwc MiddleWareChain) Append(middlewares ...MiddleWare) MiddleWareChain {
	return MiddleWareChain{middlewares: append(append(make([]MiddleWare, 0, len(mwc.middlewares)+len(middlewares)), mwc.middlewares...), middlewares...)}
}

// extend a chain with another chain
func (mwc MiddleWareChain) Extend(chain MiddleWareChain) MiddleWareChain {
	return mwc.Append(chain.middlewares...)
}