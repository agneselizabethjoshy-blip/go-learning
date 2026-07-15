package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{5, 2, 8, 1, 9}

	sort.Ints(arr)

	fmt.Println(arr)
}
