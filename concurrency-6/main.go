package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			select {
			default:
				fmt.Println("DOING SOMEWORK!")
			}
		}
	}()

	time.Sleep(10 * time.Second) // message the main to fork off back control to main from the infinite loop after 10secs
}
