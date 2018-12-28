package middleware

import (
	"io"
	"net/http"
)

//todo: Add logger

const token = "I am a token"

func main() {
	http.Handle("/", Middleware(
		http.HandlerFunc(ExampleHandler),
		AuthMiddleWare,
	))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func Middleware(h http.Handler, middleware ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middleware {
		h = mw(h)
	}
	return h
}

func AuthMiddleWare(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestToken := r.Header.Get("token")
		if len(requestToken) == 0 || requestToken != token {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnauthorized)
			io.WriteString(w, `{"error": "Invalid token"}`)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, `{"status":"ok"}`)
}
