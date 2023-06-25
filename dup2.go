package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Написать программу, которая ищет количество одинаковых строк
в последовательности файлов или, если файлов нет, в стандартном потоке ввода
*/
func main() {
	counts := make(map[string]int)
	files := os.Args[1:] // список файлов
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, errType := os.Open(arg)
			// если тип ошибки при открытии файла равен nil, то с файлом всё хорошо
			if errType != nil {
				fmt.Fprintf(os.Stderr, "dup2 has an error: %v", errType)
				continue
			}
			countLines(file, counts)
			file.Close()
		}

		for key, value := range counts {
			if value > 1 {
				fmt.Printf("line: %s, count: %d\n", key, value)
			}
		}
	}
}

func countLines(stream *os.File, counts map[string]int) {
	input := bufio.NewScanner(stream)
	for input.Scan() {
		line := input.Text()
		counts[line]++
	}
	// Мы здесь игнорируем потенциальные ошибки
}
