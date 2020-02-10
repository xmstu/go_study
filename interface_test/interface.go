// package main

// import (
// 	"fmt"
// )

// type Phone interface {
// 	call()
// }

// type NokiaPhone struct {
// }

// func (nokiaPhone NokiaPhone) call() {
// 	fmt.Println("I am Nokia, I can call you!")
// }

// type IPhone struct {
// }

// func (iPhone IPhone) call() {
// 	fmt.Println("I am iPhone, I can call you!")
// }

// func main() {
// 	var phone Phone

// 	phone = new(NokiaPhone)
// 	phone.call()

// 	phone = new(IPhone)
// 	phone.call()

// }

// package main

// import (
// 	"fmt"
// )

// type Human interface {
// 	name() string
// 	age() int
// }

// type Woman struct {
// }

// func (woman Woman) name() string {
// 	return "Jin Yawei"
// }
// func (woman Woman) age() int {
// 	return 23
// }

// type Men struct {
// }

// func (men Men) name() string {
// 	return "liweibin"
// }
// func (men Men) age() int {
// 	return 27
// }

// func main() {
// 	var man Human

// 	man = new(Woman)
// 	fmt.Println(man.name())
// 	fmt.Println(man.age())
// 	man = new(Men)
// 	fmt.Println(man.name())
// 	fmt.Println(man.age())
// }

package main

import (
	"fmt"
)

type namer interface {
	area() int
}

type rect struct {
	width, height int
}

type square struct {
	side int
}

func (r rect) area() int {
	return r.height * r.width
}

func (s square) area() int {
	return s.side * s.side
}

func main() {
	var a = rect{4, 3}
	var b = square{6}

	fmt.Println(a.area())
	fmt.Println(b.area())
}
