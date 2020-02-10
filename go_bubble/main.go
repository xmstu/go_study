package main

import "fmt"

func main() {
	arr := [...]int{21, 32, 12, 33, 34, 34, 87, 24}
	var n = len(arr)
	fmt.Println("没排序前:\n", arr)
	for i := 0; i < n-1; i++ {
		fmt.Printf("第%d次冒泡\n", i+1)
		for j := i; j <= n-1; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
			fmt.Println("排序中:\n", arr)
		}
	}
	fmt.Println("最终结果:\n", arr)

}
