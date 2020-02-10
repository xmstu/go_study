// package main

// import (
// 	. "fmt"
// )

// func main() {
// 	var numbers = make([]int, 3, 5)
// 	printSlice(numbers)
// }

// func printSlice(x []int) {
// 	Printf("len:%d cap:%d slice=%v\n", len(x), cap(x), x)
// }

// package main

// import (
// 	. "fmt"
// )

// func main() {
// 	s := []int{1, 2, 3, 10, 5, 8}
// 	Println(s[2:5])
// }

// package main

// import "fmt"

// func main() {
// 	/* 创建切片 */
// 	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
// 	printSlice(numbers)

// 	/* 打印原始切片 */
// 	fmt.Println("numbers ==", numbers)

// 	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
// 	fmt.Println("numbers[1:4] ==", numbers[1:4])

// 	/* 默认下限为 0*/
// 	fmt.Println("numbers[:3] ==", numbers[:3])

// 	/* 默认上限为 len(s)*/
// 	fmt.Println("numbers[4:] ==", numbers[4:])

// 	numbers1 := make([]int, 0, 5)
// 	printSlice(numbers1)

// 	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
// 	number2 := numbers[:2]
// 	printSlice(number2)

// 	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
// 	number3 := numbers[2:5]
// 	printSlice(number3)

// }

// func printSlice(x []int) {
// 	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
// }

package main

import "fmt"

func main() {
	//	sliceTest()
	twoDimensionArray()
}

func sliceTest() {
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:]
	for e := range s {
		fmt.Println(s[e])
	}
	s1 := make([]int, 3)
	for e := range s1 {
		fmt.Println(s1[e])
	}
}

func twoDimensionArray() {
	/* 数组 - 5 行 2 列*/
	var a = [][]int{{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
