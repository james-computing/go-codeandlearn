package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveWithCustomHandler()
	//serveWithDefaultHandler()
	//serveWithDefaultHandler2()
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	write.Write([]byte("Hello!"))
}

func serveWithCustomHandler() {
	err := http.ListenAndServe(":8080", new(HelloHandler))
	if err != nil {
		fmt.Println(err)
	}
}

func serveWithDefaultHandler() {
	http.Handle("/hello", new(HelloHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func serveWithDefaultHandler2() {
	http.Handle("/hello", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	}))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
