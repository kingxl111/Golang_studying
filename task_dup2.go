package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Написать программу, которая будет выводить имена всех файлов,
в которых есть повторяющиеся строки
*/
func main() {
	counts := make(map[string]int)
	files := os.Args[1:] // список файлов
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Отсутствуют файлы: error %d\n", 1)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%v: %v\n", arg, err)
				continue
			}

			input := bufio.NewScanner(file)
			for input.Scan() {
				line := input.Text()
				if counts[line] > 0 {
					fmt.Printf("%v\n", arg)
					break
				}
				counts[line]++
			}
			file.Close()
		}
	}
}
