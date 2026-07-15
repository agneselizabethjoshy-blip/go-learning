package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "Go is a simple and powerful language"

	words := strings.Fields(text)

	fmt.Println("Word Count:", len(words))
}
