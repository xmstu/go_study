package main

import (
	"fmt"
	"go_dev/day1/example1/goroute"
)

func main() {
	var (
		a int = 10
		b int = 12
		c     = make(chan int, 1)
	)
	go goroute.Test(a, b, c)
	fmt.Println(<-c)
}
