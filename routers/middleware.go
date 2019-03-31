package routers

import (
	"log"
	"net/http"
	"time"
)

func Logger(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	start := time.Now()
	next(w, r)
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		time.Since(start),
	)
}

func middlewareFirst(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("MiddlewareFirst - Before Handler")
	next(w, r)
	log.Println("MiddlewareFirst - After Handler")
}

func middlewareSecond(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("MiddlewareSecond - Before Handler")
	next(w, r)
	log.Println("MiddlewareSecond - After Handler")
}
