package main

// 导入系统包
import (
	"flag"
	"fmt"
)

// 定义命令行参数
var mode = flag.String("mode", "", "process mode")
var isRight = flag.Bool("is_right", false, "bool flag")

func main() {

	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode)
	fmt.Println(*isRight)
}
