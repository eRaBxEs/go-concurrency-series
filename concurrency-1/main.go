package main

import (
	"fmt"
)

func someFunc(n int) {
	fmt.Println(n)
}

func main() {
	someFunc(1)
	someFunc(2)
	someFunc(3)

	fmt.Println("Hi")
}
