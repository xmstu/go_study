package main

// package main

// import "fmt"
// import "time"

// func f1(ch chan int) {
// 	time.Sleep(time.Second * 10)
// 	ch <- 1
// }
// func f2(ch chan int) {
// 	time.Sleep(time.Second * 1)
// 	ch <- 1
// }
// func main() {
// 	var ch1 = make(chan int)
// 	var ch2 = make(chan int)
// 	go f1(ch1)
// 	go f2(ch2)
// 	select {
// 	case <-ch1:
// 		fmt.Println("The first case is selected.")
// 	case <-ch2:
// 		fmt.Println("The second case is selected.")
// 	}
// }

// package main

// import (
// 	"fmt"
// )

// //定义几个变量，其中chs和numbers分别代表通道列表和整数列表
// var ch1 chan int
// var ch2 chan int
// var chs = []chan int{ch1, ch2}
// var numbers = []int{1, 2, 3, 4, 5}

// func main() {
// 	select {
// 	case getChan(0) <- getNumber(2):
// 		fmt.Println("1th case is selected.")
// 	case getChan(1) <- getNumber(3):
// 		fmt.Println("2th case is selected.")
// 	default:
// 		fmt.Println("default!.")
// 	}
// }
// func getNumber(i int) int {
// 	fmt.Printf("numbers[%d]\n", i)
// 	return numbers[i]
// }
// func getChan(i int) chan int {
// 	fmt.Printf("chs[%d]\n", i)
// 	return chs[i]
// }

// package main

// import (
// 	"fmt"
// 	"time"
// )

// //定义几个变量，其中chs和numbers分别代表通道列表和整数列表

import (
	"fmt"
	"time"
)

var ch1 chan int = make(chan int, 1) //声明并初始化channel变量
var ch2 chan int = make(chan int, 1) //声明并初始化channel变量
var chs = []chan int{ch1, ch2}
var numbers = []int{1, 2, 3, 4, 5}

func main() {
	select {
	case getChan(0) <- getNumber(2):
		time.Sleep(time.Second * 2)
		fmt.Println("1th case is selected.")
	case getChan(1) <- getNumber(3):
		time.Sleep(time.Second * 1)
		fmt.Println("2th case is selected.")
	default:
		fmt.Println("default!.")
	}
}
func getNumber(i int) int {
	fmt.Printf("numbers[%d]\n", i)
	return numbers[i]
}
func getChan(i int) chan int {
	fmt.Printf("chs[%d]\n", i)
	return chs[i]
}

// 示例4：如果有多个case同时可以运行，go会随机选择一个case执行
//demo2.go
//package main

//import (
//	"fmt"
//)

//func main() {
//	chanCap := 5
//	ch := make(chan int, chanCap)  //创建channel，容量为5
//	for i := 0; i < chanCap; i++ { //通过for循环，向channel里填满数据
//		select { //通过select随机的向channel里追加数据
//		case ch <- 4:
//		case ch <- 2:
//		case ch <- 3:
//		case ch <- 1:
//		case ch <- 5:
//		}
//	}
//	for i := 0; i < chanCap; i++ {
//		fmt.Printf("%v\n", <-ch)
//	}
//}

// 成功: 进程退出代码 0.
// 注意：上面的案例每次运行结果都不一样。

// 示例5：使用break终止select语句的执行

// package main
// import "fmt"
// func main() {
//     var ch = make(chan int, 1)
//     ch <- 1
//     select {
//     case <-ch:
//         fmt.Println("This case is selected.")
//         break //The following statement in this case will not execute.
//         fmt.Println("After break statement")
//     default:
//         fmt.Println("This is the default case.")
//     }
//     fmt.Println("After select statement.")
// }
