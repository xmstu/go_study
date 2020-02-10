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

// package main

// import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("Mahesh", "Kumar")
// 	fmt.Println(a, b)
// }

// package main

// import "fmt"

// func main() {
// 	/* 定义局部变量 */
// 	var a int = 100
// 	var b int = 200

// 	fmt.Printf("交换前 a 的值为 : %d\n", a)
// 	fmt.Printf("交换前 b 的值为 : %d\n", b)

// 	/* 通过调用函数来交换值 */
// 	swap(a, b)

// 	fmt.Printf("交换后 a 的值 : %d\n", a)
// 	fmt.Printf("交换后 b 的值 : %d\n", b)
// }

// /* 定义相互交换值的函数 */
// func swap(x, y int) int {
// 	var temp int

// 	temp = x /* 保存 x 的值 */
// 	x = y    /* 将 y 值赋给 x */
// 	y = temp /* 将 temp 值赋给 y*/

// 	return temp
// }

// package main

// import "fmt"

// func main() {
// 	/* 定义局部变量 */
// 	var a int = 100
// 	var b int = 200

// 	fmt.Printf("交换前，a 的值 : %d\n", a)
// 	fmt.Printf("交换前，b 的值 : %d\n", b)

// 	 调用 swap() 函数
// 	 * &a 指向 a 指针，a 变量的地址
// 	 * &b 指向 b 指针，b 变量的地址

// 	swap(&a, &b)

// 	fmt.Printf("交换后，a 的值 : %d\n", a)
// 	fmt.Printf("交换后，b 的值 : %d\n", b)
// }

// func swap(x *int, y *int) {
// 	var temp int
// 	temp = *x /* 保存 x 地址上的值 */
// 	*x = *y   /* 将 y 值赋给 x */
// 	*y = temp /* 将 temp 值赋给 y */
// }

// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	/* 声明函数变量 */
// 	getSquareRoot := func(x float64) float64 {
// 		return math.Sqrt(x)
// 	}

// 	/* 使用函数 */
// 	fmt.Println(getSquareRoot(9))

// }

// package main

// import "fmt"

// func getSequence() func() int {
// 	i := 0
// 	return func() int {
// 		i += 1
// 		return i
// 	}
// }

// func main() {
// 	/* nextNumber 为一个函数，函数 i 为 0 */
// 	nextNumber := getSequence()

// 	 调用 nextNumber 函数，i 变量自增 1 并返回
// 	fmt.Println(nextNumber())
// 	fmt.Println(nextNumber())
// 	fmt.Println(nextNumber())

// 	/* 创建新的函数 nextNumber1，并查看结果 */
// 	nextNumber1 := getSequence()
// 	fmt.Println(nextNumber1())
// 	fmt.Println(nextNumber1())
// }

// package main

// import "fmt"

// func main() {
// 	add_func := add(1, 2)
// 	fmt.Println(add_func())
// 	fmt.Println(add_func())
// 	fmt.Println(add_func())
// }

// // 闭包使用方法
// func add(x1, x2 int) func() (int, int) {
// 	i := 0
// 	return func() (int, int) {
// 		i++
// 		return i, x1 + x2
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	add_func := add(1, 2)
// 	fmt.Println(add_func(1, 1))
// 	fmt.Println(add_func(0, 0))
// 	fmt.Println(add_func(2, 2))
// }

// // 闭包使用方法
// func add(x1, x2 int) func(x3 int, x4 int) (int, int, int) {
// 	i := 0
// 	return func(x3 int, x4 int) (int, int, int) {
// 		i++
// 		return i, x1 + x2, x3 + x4
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// /* 定义结构体 */
// type Circle struct {
// 	radius float64
// }

// func main() {
// 	var c1 Circle
// 	c1.radius = 10.00
// 	fmt.Println("Area of Circle(c1) = ", c1.getArea())
// }

// //该 method 属于 Circle 类型对象中的方法
// func (c Circle) getArea() float64 {
// 	//c.radius 即为 Circle 类型对象中的属性
// 	return 3.14 * c.radius * c.radius
// }

package main

import "fmt"

//行数
const LINES int = 10

// 杨辉三角
func ShowYangHuiTriangle() {
	nums := []int{}
	for i := 0; i < LINES; i++ {
		//补空白
		for j := 0; j < (LINES - i); j++ {
			fmt.Print(" ")
		}
		for j := 0; j < (i + 1); j++ {
			var length = len(nums)
			var value int
			if j == 0 || j == i {
				value = 1
			} else {
				value = nums[length-i] + nums[length-i-1]
			}
			nums = append(nums, value)
			fmt.Print(value, " ")
		}
		fmt.Println("")
	}
}

func main() {
	ShowYangHuiTriangle()
}
