// package main

// import (
// 	. "fmt"
// )

// /* 函数返回两个数的最大值 */
// func max(num1, num2 int) int {
// 	/* 声明局部变量 */
// 	var result int

// 	if num1 > num2 {
// 		result = num1
// 	} else {
// 		result = num2
// 	}
// 	return result
// }

// func main() {
// 	var a, b int = 10, 5
// 	var ret int
// 	ret = max(a, b)
// 	Printf("最大值是%d", ret)
// }

package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Mahesh", "Kumar")
	fmt.Println(a, b)
}
