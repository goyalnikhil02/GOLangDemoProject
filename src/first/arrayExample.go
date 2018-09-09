package main

import "fmt"

func main() {
	var x [2]int
	x[1] = 1
	x[0] = 2
	//x[3]=8
	fmt.Print("Hi", x[1]+x[0])

}
