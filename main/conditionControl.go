package main

import (
	"fmt"
	"log"
)

func condition() {
	num := 10
	if num > 10 {
		fmt.Println(">", num)
		log.Fatal(">", num)
	} else {
		fmt.Println("not > ", num)
		log.Fatal("not > ", num)
	}
	fmt.Println("会运行吗")
}
