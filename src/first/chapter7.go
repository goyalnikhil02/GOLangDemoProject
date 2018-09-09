package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Print("Enter 2  numbers: ")
	var input1 int8
	var input2 int8
	fmt.Scanf("%d", &input1)
	fmt.Scanf("%d", &input2)

	//fmt.Println("You Entereed :", input1)

	fmt.Println(average(input1, input2))
}
func average(x int8, y int8) int8 {
	str1 := "1234"

	/** converting the str1 variable into an int using Atoi method */
	i1, err := strconv.Atoi(str1)
	if err == nil {
		fmt.Println("conversion:", i1)
	}

	str2 := "102"

	/** converting the str1 variable into an int using Atoi method */
	i2, err := strconv.ParseInt(str2, 10, 8)
	if err == nil {
		fmt.Println("value is:", i2)
	} else {
		fmt.Println("error in conversion")
	}

	fmt.Println("y:", y)
	return x + y
}
