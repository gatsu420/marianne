package middlewares

import "net/http"

type middleware func(http.Handler) http.Handler

func Stack(h http.Handler, mw ...middleware) http.Handler {
	for i := len(mw) - 1; i >= 0; i-- {
		h = mw[i](h)
	}
	return h
}
