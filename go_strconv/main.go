package main

import (
	"fmt"
	"strconv"
)

func parseFunc() {
	b, _ := strconv.ParseBool("true")
	f, _ := strconv.ParseFloat("3.1415", 64)
	l, _ := strconv.ParseInt("-42", 10, 64)
	u, _ := strconv.ParseUint("42", 10, 64)
	fmt.Println("b:", b)
	fmt.Println("f:", f)
	fmt.Println("l:", l)
	fmt.Println("u:", u)

}

func main() {
	i, _ := strconv.Atoi("3")
	fmt.Println(3 + i) // 6
	// Atoi()转换失败
	i, err := strconv.Atoi("a")
	if err != nil {
		fmt.Println("converted failed")
	}

	fmt.Println("a" + strconv.Itoa(32))
	parseFunc()
}
