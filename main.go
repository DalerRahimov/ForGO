package main

import (
	"HW7/Funcs"
	"fmt"
)

func main() {

	fmt.Println("Выберите задачу: 1)if28 2)for39 3)for40")
	var ex int

	switch ex {
	case 1:
		Funcs.For38()
	case 2:
		Funcs.For39()
	case 3:
		Funcs.For40()
	default:
		fmt.Println("Error")

	}
}
