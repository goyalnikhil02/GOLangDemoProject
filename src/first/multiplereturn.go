package main

import (
	"fmt"
	"strconv"
)

func main() {

	x, y := reurnTWO("12")
	fmt.Print("x:", x)
	fmt.Println("y:", y)
}

func reurnTWO(a string) (int64, error) {

	/** converting the str1 variable into an int using Atoi method */
	i2, err := strconv.ParseInt(a, 10, 8)
	if err == nil {
		return i2, err
		//		fmt.Println("value is:",i2)
	} else {
		fmt.Println("error in conversion")
	}

}
