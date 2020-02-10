package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name:")
	// 读取数据直到\n为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occured:%s\n", err)
		// 异常退出
		os.Exit(1)
	} else {
		// 用切片操作删除最后的\n
		name := input[:len(input)-2]
		fmt.Printf("hello, %s!What can I do for you?\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An err occurred:%s\n", err)
			continue
		}

		input = input[:len(input)-2]
		// 全部转换成小写
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry,I didn't catch you")

		}
	}

}
