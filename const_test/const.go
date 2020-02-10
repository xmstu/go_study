package main

import "fmt"

const (
	i = 1 << iota
	j = 3 << iota
	k = 4 << iota
	l = 4 << iota
)

func main() {
	fmt.Println("i=", i)
	fmt.Println("j=", j)
	fmt.Println("k=", k)
	fmt.Println("l=", l)
}
