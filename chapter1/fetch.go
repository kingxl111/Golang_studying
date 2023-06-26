package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//Выборка содержимого URL
//Напишем программу, которая будет делать http запрос по указанному url, а также выводить код состояния HTTP
func main() {
	//Если в url нет http://, программа будет его дописывать
	for _, cur_url := range os.Args[1:] {
		if !strings.HasPrefix(cur_url, "http://") {
			cur_url = "http://" + cur_url
		}
		resp, err := http.Get(cur_url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: error %v\n", err)
			os.Exit(1)
		}

		_, er := io.Copy(os.Stdout, resp.Body) // в качестве _ возвращает количество скопированных байтов
		fmt.Fprintf(os.Stdout, "Response status code: %v", resp.Status) //Статус ответа
		resp.Body.Close()
		if er != nil {
			fmt.Fprintf(os.Stderr, "fetch: error %v\n", er)
			os.Exit(1)
		}
		// fmt.Fprintf(os.Stdout, "%s\n", str)
	}
}
