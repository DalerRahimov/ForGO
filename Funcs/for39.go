package Funcs

import (
	"fmt"
	"log"
)

func For39() {
	fmt.Println("Введите 2 целых числа")
	var a, b int
	var err error

	_, err = fmt.Scan(&a)
	if err != nil {
		log.Println("Input error")
		return
	}
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Println("Input error")
		return
	}

	if a > b {
		fmt.Printf("Input error")
		return
	} else {
		for i := a; i <= b; i++ {
			for x := i; x > 0; x-- {
				fmt.Printf("%v ", i)
			}
			fmt.Println("")
		}
	}

}
