package main

import "fmt"

func nineNine() {
	for y := 1; y <= 9; y++ {
		for x := 1; x <= y; x++ {
			fmt.Printf("%d*%d=%d ", x, y, x*y)
		}
		fmt.Println()
	}
}

func forChannel() {
	c := make(chan int)
	go func() {
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

func main() {
	var i int
	// 类似于while i<= 10
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	nineNine()
	forChannel()

}
