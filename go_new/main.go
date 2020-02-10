package main

import "fmt"

func main() {
	// new() 函数可以创建一个对应类型的指针，创建过程会分配内存，被创建的指针指向默认值。
	str := new(string)
	*str = "go教程"
	fmt.Printf("ptr str: %p", str)
	fmt.Printf("\nvalue str: %v", *str)

}
