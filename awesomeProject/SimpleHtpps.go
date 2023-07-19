package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
		<html>
			<head>
				<title>Hello World</title>
			</head>
		<body>
			Hello World!
		</body>
		</html>`,
	)
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi\n My name is\n Mobarak")
}

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, strconv.Itoa(counter))
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter += 20
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func decrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter -= 20
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//		//fmt.Println(r.URL.Path)
//		http.ServeFile(w, r, "")
//	})
//	http.HandleFunc("/increment", incrementCounter)
//	http.HandleFunc("/decrement", decrementCounter)
//	http.HandleFunc("/hello", hello)
//	http.HandleFunc("/hi", hi)
//	http.ListenAndServe(":9000", nil)
//}
