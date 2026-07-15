package main

import "fmt"

func main() {
	arr := []int{10, 25, 5, 99, 34}
	max := arr[0]

	for _, value := range arr {
		if value > max {
			max = value
		}
	}

	fmt.Println("Maximum:", max)
}
