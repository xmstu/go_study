package main

import "fmt"

// break 语句可以结束 for、switch 和 select 的代码块，另外 break 语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的 for、switch 和 select 的代码块上。
func main() {
OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				break OuterLoop
			case 3:
				fmt.Println(i, j)
				break OuterLoop

			}
		}
	}

}
