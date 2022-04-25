package main

import "fmt"

func variable() {
	a, b := 20, 30
	fmt.Println("a is", a, "b is ", b)
	c := true
	d := false
	fmt.Println("c is ", c, "d is", d)

	var e = 5.9 / 8
	fmt.Printf("a's type %T value %v ", e, e)
}
