package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var req_count int

func main() {
	server := "localhost:8000"
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)          // Если дополнить адрес сервера с помощью /count, вызовется counter
	log.Fatal(http.ListenAndServe(server, nil)) //Запускаем наш сервер
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Requests counter: %d\n", req_count)
	mu.Unlock()
}

func handler(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	req_count++
	mu.Unlock()
	fmt.Fprintf(writer, "Request has been processed!\n URL path: %v\n", request.URL.Path)
}
