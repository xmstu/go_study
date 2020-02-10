// package main

// import "fmt"

// func main() {
// 	/* 定义局部变量 */
// 	var grade string = "B"
// 	var marks int = 60

// 	switch marks {
// 	case 90:
// 		grade = "A"
// 	case 80:
// 		grade = "B"
// 	case 60, 70:
// 		grade = "C"
// 	case 50:
// 		grade = "F"
// 	default:
// 		grade = "D"
// 	}

// 	switch {
// 	case grade == "A":
// 		fmt.Printf("优秀!\n")
// 	case grade == "B", grade == "C":
// 		fmt.Printf("良好\n")
// 	case grade == "D":
// 		fmt.Printf("及格\n")
// 	case grade == "F":
// 		fmt.Printf("不及格\n")
// 	default:
// 		fmt.Printf("差\n")
// 	}
// 	fmt.Printf("你的等级是 %s\n", grade)
// }

package main

import "fmt"

func main() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
