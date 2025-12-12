package middleware

import (
	"fmt"
	"net/http"
)

func Recover() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Print("Recover -> ")
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(500)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
