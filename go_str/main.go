package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"strings"
	"unicode/utf8"
)

func utf8Count() {
	fmt.Println(utf8.RuneCountInString("忍者"))
	fmt.Println(utf8.RuneCountInString("龙忍出鞘,fight!"))

}

func findSubstr() {
	tracer := "死神来了, 死神bye bye"
	comma := strings.Index(tracer, ", ")

	pos := strings.Index(tracer[comma:], "死神")

	fmt.Println(comma, pos, tracer[comma+pos:])
}

func strChange() {
	angel := "Heros never die"

	angleBytes := []byte(angel)

	for i := 5; i <= 10; i++ {
		angleBytes[i] = ' '
	}

	fmt.Println(string(angleBytes))

}

func strBuild() {
	hammer := "吃我一锤"
	sickle := "死吧"
	// 声明字节缓冲
	var stringBuilder bytes.Buffer
	// 把字符串写入缓冲
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sickle)
	// 将缓冲以字符串形式输出
	fmt.Println(stringBuilder.String())

}

func strFormat() {
	var progress = 2
	var target = 8

	// 两参数格式化
	title := fmt.Sprintf("已采集%d个药草, 还需要%d个完成任务", progress, target)

	fmt.Println(title)

	pi := 3.14159
	// 按数值本身的格式输出
	variant := fmt.Sprintf("%v %v %v", "月球基地", pi, true)

	fmt.Println(variant)

	// 匿名结构体声明, 并赋予初值
	profile := &struct {
		Name string
		HP   int
	}{
		Name: "rat",
		HP:   150,
	}

	fmt.Printf("使用'%%+v' %+v\n", profile)

	fmt.Printf("使用'%%#v' %#v\n", profile)

	fmt.Printf("使用'%%T' %T\n", profile)
}

func strEncode() {
	// 需要处理的字符串
	message := "Away from keyboard. https://golang.org/"
	// 编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	// 输出编码完成的消息
	fmt.Println(encodedMessage)
	// 解码消息
	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	// 出错处理
	if err != nil {
		fmt.Println(err)
	} else {
		// 打印解码完成的数据
		fmt.Println(string(data))
	}

}

func strUnicode() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'
	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point

}

func main() {
	str := "hello " +
		"world"
	fmt.Println(str)
	const str1 = `
	第一行
	第二行
	第三行
	\r\n
	`
	fmt.Println(str1)

	utf8Count()

	// 遍历ascii字符
	theme := "狙击 start"

	for i := 0; i < len(theme); i++ {
		fmt.Printf("ascii: %c %d\n", theme[i], theme[i])
	}
	// 遍历Unicode字符
	for _, s := range theme {
		fmt.Printf("Unicode: %c, %d\n", s, s)
	}

	findSubstr()
	strChange()
	strBuild()
	strFormat()
	strEncode()
	strUnicode()

}
