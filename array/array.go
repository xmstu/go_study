// package main

// import (
// 	. "fmt"
// )

// func main() {
// 	var n [10]int /* n是一个长度为10的数组 */
// 	var i, j int

// 	for i = 0; i < 10; i++ {
// 		n[i] = i + 100
// 	}

// 	for j = 0; j < 10; j++ {
// 		Printf("[%d]=%d\n", j, n[j])
// 	}
// }

// package main

// import (
// 	. "fmt"
// )

// func main() {
// 	var a = [5][2]int{{0, 0}, {1, 2}, {2, 4}, {3, 6}, {4, 8}}
// 	var i, j int

// 	for i = 0; i < 5; i++ {
// 		for j = 0; j < 2; j++ {
// 			Printf("a[%d][%d] = %d\n", i, j, a[i][j])
// 		}
// 	}
// }

// package main

// import (
// 	. "fmt"
// )

// func main() {
// 	var balance = []float32{100.15, 20.0, 30.6, 17.9, 50.8}
// 	var avg float32

// 	avg = getAverage(balance, len(balance))
// 	Printf("平均值为:%f", avg)
// }

// func getAverage(arr []float32, size int) float32 {
// 	var i int
// 	var avg, sum float32

// 	for i = 0; i < size; i++ {
// 		sum += arr[i]
// 	}
// 	avg = float32(sum) / float32(size)
// 	return avg
// }

// package main

// import "fmt"

// const MAX int = 3

// func main() {
// 	a := []int{10, 100, 200}
// 	var i int
// 	var ptr [MAX]*int

// 	for i = 0; i < MAX; i++ {
// 		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
// 	}

// 	for i = 0; i < MAX; i++ {
// 		fmt.Printf("a[%d] = %d\n", i, ptr[i])
// 		fmt.Printf("a[%d] = %d\n", i, *ptr[i])
// 	}
// }

package main

import "fmt"

func main() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("a = %d\n", a)
	fmt.Printf("*ptr = %d\n", *ptr)
	fmt.Printf("ptr = %d\n", ptr)
	fmt.Printf("**pptr = %d\n", **pptr)
	fmt.Printf("*pptr = %d\n", *pptr)
	fmt.Printf("pptr = %d\n", pptr)
}
