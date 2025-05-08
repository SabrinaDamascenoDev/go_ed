
package main

import "fmt"

func eh_primo(x int, div int) bool {
	if div == 1 {
		return true
	}
	if x%div == 0 {
		return false
	}
	return eh_primo(x, div-1)
} 

func main() {
	var x int
	fmt.Scan(&x)
	if x < 2 {
		fmt.Println("false")
	} else if eh_primo(x, x-1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
