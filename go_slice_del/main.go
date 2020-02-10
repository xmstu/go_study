package main

import "fmt"

func delHead() {
	// 移动数据指针删除开头元素
	a := []int{1, 2, 3}
	a = a[1:]
	fmt.Println("a:", a)

}

func delAppend() {
	// 也可以不移动数据指针，但是将后面的数据向开头移动，可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）
	a := []int{1, 2, 3}
	fmt.Println("new a:", a[:0])
	fmt.Println("new a:", a[1:])
	a = append(a[:0], a[1:]...)
	fmt.Println("a:", a)
}

func delCopy() {
	// 通过copy函数来删除开头的元素
	a := []int{1, 2, 3}
	a = a[:copy(a, a[1:])]
	fmt.Println("a:", a)
}

func delMiddle() {
	// 从中间位置删除元素
	const count = 10
	a := make([]int, count)
	for i := 0; i < count; i++ {
		a[i] = i
	}
	fmt.Println("a:", a)
	a = append(a[:5], a[5+1:]...)
	fmt.Println("a:", a)
}

func delTail() {
	// 从尾部删除元素
	a := []int{1, 2, 3}
	a = a[:len(a)-1]
	fmt.Println("a:", a)

}

func main() {
	delHead()
	delAppend()
	delCopy()
	delMiddle()
	delTail()

}
