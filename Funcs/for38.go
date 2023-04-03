package Funcs

import (
	"fmt"
	"log"
)

func For38() {

	fmt.Println("Введите число")
	var n int
	res := 0
	var err error
	_, err = fmt.Scan(&n)
	if err != nil {
		log.Println("Input error")
		return
	}

	for i := 1; i <= n; i++ {
		ni := i
		for j := 1; j <= n-i; j++ {
			ni = ni * i
			res = res + ni
		}
	}
	fmt.Println(res)
}
