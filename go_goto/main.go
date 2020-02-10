package main

import "fmt"

func main() {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 终止循环，跳到标签
				fmt.Printf("y:%d\n", y)
				goto breakHere
			}
		}
	}

	return
	// 标签
breakHere:
	fmt.Println("done")
}
