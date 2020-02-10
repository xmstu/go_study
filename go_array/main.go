package main

import "fmt"

func main() {
	var array [4][2]int
	fmt.Println("array:", array)
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	fmt.Println("array:", array)
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	fmt.Println("array:", array)
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println("array:", array)
}
