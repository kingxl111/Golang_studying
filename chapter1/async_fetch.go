package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

//Задача: написать горутину, которая будет обрабатывать сразу несколько URL, поданных на вход
//В итоге результат работы должен быть сохранен в файле
func main() {
	start := time.Now()
	ch := make(chan string)

	for _, cur_url := range os.Args[1:] {
		go fetch(cur_url, ch)
	}

	file_name := "async_fetch_result.txt"
	file, e := os.Create(file_name)
	if e != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", e)
		os.Exit(1)
	}
	for range os.Args[1:] {
		// fmt.Println(<-ch)
		fmt.Fprintln(file, <-ch) // Шикарные автоматические переносы строки
	}
	// fmt.Fprintf(os.Stdout, "Total: %.1f\n", time.Since(start).Seconds())
	fmt.Fprintf(file, "Total: %.1f\n", time.Since(start).Seconds())
	file.Close()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Sprintf возвращает строку, а не печатает в какой-либо поток
		return
	}

	number_bytes, er := io.Copy(ioutil.Discard, resp.Body)
	if er != nil {
		ch <- fmt.Sprint(err)
		return
	}
	resp.Body.Close()

	seconds := time.Since(start).Seconds()
	ch <- fmt.Sprintf("time: %.2f, nbytes: %d, URL: %s", seconds, number_bytes, url)
}
