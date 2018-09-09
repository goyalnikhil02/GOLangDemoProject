package main

import "fmt"

var stringData string = "Test Scope"

func main() {
	test2()
	var data string = "My hello string"

	for _, r := range data {
		c := string(r)
		fmt.Print(c)
	}
	fmt.Println("###############")
	fmt.Println("Finding the Length of String:", len("Hello Nikhil"))
	c := string("Hello Nikhil"[1])
	fmt.Println(c)
	fmt.Println("Hello" + "World")
}

func test2() {
	test3()
	//inistailizing any variable
	x := "demo"
	var y string = "dem" + "o"
	fmt.Println(x == y)
}

func test3() {
	fmt.Println("Scope" + stringData)
}
