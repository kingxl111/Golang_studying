package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		var line = input.Text()
		counts[line]++
	}

	for key, value := range counts { // так как это отображение(map), здесь берется рандомный элемент
		if value > 1 {
			fmt.Printf("count: %d, string: %s\n", value, key)
		}

	}
}
