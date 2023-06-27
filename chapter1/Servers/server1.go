package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := "localhost:8000"
	http.HandleFunc("/", handler) // Если производится запрос, то вызывается обработчик этого запроса
	log.Fatal(http.ListenAndServe(server, nil))
}

func handler(wrtr http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(wrtr, "Cur URL.PATH: %v\n", req.URL.Path) // На сервере печатаем путь
}

/* 
Можно еще проделать вот такую штуку: 
1) Запускаем свой сервер в фоне
2) Делаем Get запрос на этот сервер с помощью нашей программки fetch.go
3) Получаем результат в командной строке

$ go build server1.go
$ ./server1 &
$ go build fetch.go
$ ./fetch http://localhost:8000/dfjaldsmkslade
$ 

vadim@vadim-GF65-Thin-9SEXR:~/Go/Golang_studying/chapter1$ go build server1.go 
vadim@vadim-GF65-Thin-9SEXR:~/Go/Golang_studying/chapter1$ ./server1 &
[1] 7166
vadim@vadim-GF65-Thin-9SEXR:~/Go/Golang_studying/chapter1$ go build fetch.go 
vadim@vadim-GF65-Thin-9SEXR:~/Go/Golang_studying/chapter1$ ./fetch http://localhost:8000/dfjaldsmkslade
Cur URL.PATH: /dfjaldsmkslade
Response status code: 200 OKvadim@vadim-GF65-Thin-9SEXR:~/Go/Golang_studying/chapter1$
*/
