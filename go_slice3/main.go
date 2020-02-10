package main

import "fmt"

func sliceAppend() {
	var numbers []int

	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d cap: %d pointer: %p\n", len(numbers), cap(numbers), numbers)
	}

}

func sliceHeadAppend() {
	var a = []int{1, 2, 3}
	a = append([]int{0}, a...)
	fmt.Println("a:", a)
	fmt.Printf("len: %d cap: %d pointer: %p\n", len(a), cap(a), a)
	a = append([]int{-3, -2, -1}, a...)
	fmt.Println("a:", a)
	fmt.Printf("len: %d cap: %d pointer: %p\n", len(a), cap(a), a)
}

func sliceMultiAppend() {
	// 每个添加操作中的第二个 append 调用都会创建一个临时切片，并将 a[i:] 的内容复制到新创建的切片中，然后将临时创建的切片再追加到 a[:i] 中。
	var a = []int{7,8,9}
	a = append(a[:1], append([]int{6}, a[1:]...)...)
	fmt.Println("a:", a)
	a = append(a[:1], append([]int{1,2,3}, a[1:]...)...)
	fmt.Println("a:", a)
}

func main() {
	a := make([]int, 2)
	b := make([]int, 2, 10)
	fmt.Println(a, b)
	fmt.Println(len(a), len(b))

	sliceAppend()
	sliceHeadAppend()
	sliceMultiAppend()
}
