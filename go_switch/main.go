package main

import "fmt"

func main() {
	var a = "hello"
	switch a {
	case "hello":
		fmt.Println(1)
	case "world":
		fmt.Println(2)
	default:
		fmt.Println(0)

	}

	var b = "dad"
	switch b {
	case "mum", "dad":
		fmt.Println("famliy")

	}

	var r int = 11
	switch {
	case r > 10 && r < 20:
		fmt.Println(r)

	}

}
