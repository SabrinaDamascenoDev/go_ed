package main

import (
	"fmt"
)

func mdc(a, b int) int {
	if a == 0{
		return b
	}
	if b == 0{
		return a
	}
	if a%b == 0 {
		return b
	}
	resto := a%b
	a = b
	b = resto
	b = mdc(a, b)
	return b
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
