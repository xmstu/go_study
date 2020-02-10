package main

import "fmt"

// 代码说明：第 14 行将结束当前循环，开启下一次的外层循环，而不是第 10 行的循环。
func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				continue OuterLoop
			}
		}
	}

}
