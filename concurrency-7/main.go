package main

import (
	"fmt"
	"time"
)

func doWork(done <-chan bool) { // read only channel;sender kind of channel
	for {
		select {
		case <-done: // receive from the done channel; to enable the parent goroutine cancel the default task
			return
		default:
			fmt.Println("DOING WORK")
		}
	}
}

func main() {

	done := make(chan bool)

	go doWork(done)

	time.Sleep(3 * time.Second)

	close(done) // With these statement the parent goroutine sends a message

}
