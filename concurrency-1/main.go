package main

import (
	"fmt"
)

func someFunc(n int) {
	fmt.Println(n)
}

func main() {
	// with go routine added to each function, they are said to run not in sync with main, they are therefore asynchronous
	go someFunc(1)
	go someFunc(2)
	go someFunc(3)

	fmt.Println("Hi")
}
