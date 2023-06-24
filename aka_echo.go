package main

import (
	"fmt"
	"os"
)

// Напишем аналог утилиты echo
func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ { //Начинаем с 1, так как Args[0] - это имя самой команды
		s = s + sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
