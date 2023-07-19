package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// HTTP middleware setting a value on the request context
func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create new context from `r` request context, and assign key `"user"`
		// to value of `"123"`
		ctx := context.WithValue(r.Context(), "user", "123")

		// call the next handler in the chain, passing the response writer and
		// the updated request object with the new context value.
		//
		// note: context.Context values are nested, so any previously set
		// values will be accessible as well, and the new `"user"` key
		// will be accessible from this point forward.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	r := chi.NewRouter()
	//r.Use(middleware.RealIP)
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This is homepage")
	})

	//r.Get("/api/hello", func(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprintf(w, "Hello there!")
	//})

	r.Get("/{date}-{slug}", func(w http.ResponseWriter, r *http.Request) {
		dateParam := chi.URLParam(r, "date")
		slugParam := chi.URLParam(r, "slug")
		fmt.Fprintf(w, "This is %s and %s", dateParam, slugParam)
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte("route does not exist"))
	})
	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(405)
		w.Write([]byte("method is not valid"))
	})

	apiRouter := chi.NewRouter()
	apiRouter.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello this is subrouter")
	})
	r.Mount("/api", apiRouter)
	http.ListenAndServe(":3000", r)
}
