package main

import (
    "fmt"
    "strconv"
)

func main() {
    // 声明一个slice
    b10 := []byte("int (base 10):")
   
    // 将转换为10进制的string，追加到slice中
    b10 = strconv.AppendInt(b10, -42, 10)
    fmt.Println(string(b10))

    b16 := []byte("int (base 16):")
    b16 = strconv.AppendInt(b16, -42, 16)
    fmt.Println(string(b16))
}