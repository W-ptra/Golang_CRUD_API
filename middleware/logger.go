package middleware

import (
	"log"
	"net/http"
)

func Logger(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request){
		log.Printf("%v %v %v %v\n",r.Method,r.Host,r.URL.Path,r.Body)
		next.ServeHTTP(w,r)
	}
}