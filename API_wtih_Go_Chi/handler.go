package main

import (
	"encoding/json"
	"net/http"
)

func GetAvailableBookId() int {
	for i := 1; ; i++ {
		_, ok := BookList[i]
		if !ok {
			return i
		}
	}
}

func (book *Book) AddBook() {
	i := GetAvailableBookId()
	book.BookId = i
	AuthorList[book.AuthorId].MyBooks[i] = true
	AuthorList[book.AuthorId].TotalBook++
	BookList[i] = book
	//x, _ := json.Marshal(book)
	//fmt.Printf("%+v\n", string(x))
	//fmt.Printf("%+v", AuthorList[1])
}

func HandleAddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	if _, ok := AuthorList[newBook.AuthorId]; ok != true {
		w.WriteHeader(http.StatusForbidden)
		http.Error(w, "This author is not available first add the author!", 400)
		return
	}
	newBook.AddBook()
	err = json.NewEncoder(w).Encode(BookList)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
