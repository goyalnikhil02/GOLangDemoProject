package main

import (
	"fmt"
	"runtime"
)

const d = 12

var (
	a = 10
	b = 11
	c = 13
)

func main() {
	fmt.Println("Sum will be :", (a + b + c))
	fmt.Print("Enter a number: ")
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println("You Entereed :", input)
	fmt.Println("Value will be :", d+1)

	switchExample(input)
}

func forExam() {

	i := 1

	for i < 10 {
		fmt.Println(i)
		i += 1
	}
	fmt.Println("#################")
	for a := 1; a <= 5; a++ {
		if a%2 == 0 {
			fmt.Println("value :", a)
		}

	}
}

func switchExample(i string) {
	switch i {
	case "ONE":
		fmt.Println("ONE")
	case "TWO":
		fmt.Println("TWO")
	default:
		fmt.Println("Enter correct value")
	}

	unserInput()

}

func unserInput() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}


