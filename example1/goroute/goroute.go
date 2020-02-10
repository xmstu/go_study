package goroute

// 被外部调用的方法必须大写首字母
func Test(a, b int, c chan int) {
	c <- a + b
}
