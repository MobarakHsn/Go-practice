package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	Init()
	GenerateBook()
	GenerateAuthor()
	GenerateOrder()
	GenerateCustomer()
	//var book = Book{
	//	BookId:       1,
	//	Title:        "AN INTRODUCTION TO PROGRAMMING IN GO",
	//	Language:     "English",
	//	NumberOfPage: 161,
	//	Price:        200,
	//	AuthorId:     1,
	//}
	//book.AddBook()
	//fmt.Printf("%#v", CustomerList[1])
	//fmt.Printf("%#v", BookList[1])

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/Book", HandleAddBook)
	http.ListenAndServe(":8081", r)
}

//func main() {
//	r := chi.NewRouter()
//	//r.Use(middleware.RealIP)
//	r.Use(middleware.Logger)
//
//	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "This is homepage")
//	})
//
//	//r.Get("/api/hello", func(w http.ResponseWriter, r *http.Request) {
//	//	fmt.Fprintf(w, "Hello there!")
//	//})
//
//	r.Get("/{date}-{slug}", func(w http.ResponseWriter, r *http.Request) {
//		dateParam := chi.URLParam(r, "date")
//		slugParam := chi.URLParam(r, "slug")
//		fmt.Fprintf(w, "This is %s and %s", dateParam, slugParam)
//	})
//	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(404)
//		w.Write([]byte("route does not exist"))
//	})
//	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
//		w.WriteHeader(405)
//		w.Write([]byte("method is not valid"))
//	})
//
//	apiRouter := chi.NewRouter()
//	apiRouter.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Fprintf(w, "Hello this is subrouter")
//	})
//	r.Mount("/api", apiRouter)
//	http.ListenAndServe(":3000", r)
//}
