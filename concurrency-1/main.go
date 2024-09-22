package main

import (
	"fmt"
	"time"
)

func someFunc(n int) {
	fmt.Println(n)
}

func main() {
	// with go routine added to each function, they are said to run not in sync with main, they are therefore asynchronous
	go someFunc(1)
	go someFunc(2)
	go someFunc(3)

	// inorder to make main wait for them to be executed we add a delay
	time.Sleep(time.Second * 2)
	// the order by which they are printed shows that they are all asynchronous, whichever is faster is printed first/
	// they are all running concurrently

	fmt.Println("Hi")
}
