package main

import (
	"fmt"
	"runtime"
	"time"
)

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

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}

	// 没有条件的switch同switch true一样，这种形式能将一长串if-then-else写得更加清晰
	t := time.Now()
	switch {
	case t.Hour() >= 6 && t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() >= 12 && t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}
