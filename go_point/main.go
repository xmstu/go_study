package main

import (
	"fmt"
)

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func swap2(a, b *int) {
	fmt.Printf("\nptr a: %p, b: %p", a, b)
	b, a = a, b
	fmt.Printf("\nptr a: %p, b: %p", a, b)

}

func main() {

	// 准备一个字符串类型
	var house = "Malibu Point 10880, 90265"

	// 对字符串取地址, ptr类型为*string
	ptr := &house

	// 打印ptr的类型
	fmt.Printf("ptr type: %T\n", ptr)

	// 打印ptr的指针地址
	fmt.Printf("address: %p\n", ptr)

	// 对指针进行取值操作
	value := *ptr

	// 取值后的类型
	fmt.Printf("value type: %T\n", value)

	// 指针取值后就是指向变量的值
	fmt.Printf("value: %s\n", value)

	a, b := 10, 20
	swap(&a, &b)
	fmt.Printf("a: %d, b: %d", a, b)
	swap2(&a, &b)
	fmt.Printf("\na: %d, b: %d", a, b)

}
