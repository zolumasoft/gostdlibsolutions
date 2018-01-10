package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var requestsServed uint64

func greetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
	log.Println("GREETED")
}

type statsHandler struct{}

func (sh *statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Requests Handled: %d\n", atomic.LoadUint64(&requestsServed))
	log.Println("STATS PROVIDED")
}

func counter(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		atomic.AddUint64(&requestsServed, 1)
		log.Println("COUNTER >> Counted")
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("LOGGER >> START %s %q\n", r.Method, r.URL.String())
		t := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("LOGGER >> END %s %q (%v)\n", r.Method, r.URL.String(), time.Now().Sub(t))
	})
}

func use(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

func main() {
	http.Handle("/greet", use(http.HandlerFunc(greetHandler), counter, logger))

	sh := &statsHandler{}
	http.Handle("/stats", use(sh, logger))

	log.Println("Starting server...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8090", nil))
}

// Two styles of middleware because of the following:
// http.Handle("/", http.HandleFunc(f)) equivalent to
// http.HandleFunc("/", f)
// if f has signature func(http.ResponseWriter, *http.Request)

func middlewareUsingHandlerFunc(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// the middleware's logic here
		f(w, r) // equivalent to f.ServeHTTP(w, r)
	}
}

func middlewareUsingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// the middleware's logic here
		next.ServeHTTP(w, r)
	})
}
