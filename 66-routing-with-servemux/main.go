package main

import (
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello!"))
	})

	router.HandleFunc("POST /hello/{user_id}", func(writer http.ResponseWriter, request *http.Request) {
		userId := request.PathValue("user_id")
		writer.Write([]byte(fmt.Sprintf("hello %s", userId)))
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}
