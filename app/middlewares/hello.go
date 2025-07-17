package middlewares

import (
	"fmt"
	"net/http"
)

func Hello(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hello")
		h.ServeHTTP(w, r)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (r *statusRecorder) WriteHeader(code int) {
	r.ResponseWriter.WriteHeader(code)
	r.statusCode = code
}

func Ngetes(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := &statusRecorder{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		h.ServeHTTP(rec, r)
		fmt.Printf("%v", rec.statusCode)
		fmt.Println("Ngetes")
	})
}
