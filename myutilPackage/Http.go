package myutil

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

var authenticatedUsers = map[string]struct{}{
	"John":  {},
	"Susan": {},
}

func handlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Println("Start", t)
		fmt.Fprintln(w, "Hi!")
		next.ServeHTTP(w, r)
		log.Println("End", time.Since(t))
	})
}

func authenticateUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.URL.Query().Get("user")
		log.Println(time.Now(), user)
		if _, ok := authenticatedUsers[user]; !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), "user", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func greetUser(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	fmt.Fprintln(w, "Hello", user)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!\n", r.Header.Get("User-Agent"))
}

func Http() {
	http.HandleFunc("/", handler)
	http.Handle("/api", handlerMiddleware(http.HandlerFunc(handler)))
	http.Handle("/login", authenticateUser(http.HandlerFunc(greetUser)))
	fmt.Println("Starting server at port: 8080")
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		fmt.Println(e)
	}
}
