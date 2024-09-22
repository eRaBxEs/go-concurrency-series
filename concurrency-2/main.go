package main

import (
	"fmt"
)

func main() {

	myChannel := make(chan string)

	go func() {
		myChannel <- "data" // sending a data to the declared channel
	}()

	msg := <-myChannel // this line of code here is blocking since it is receiving from a channel
	// the main channel either waits for this to be closed or for the message to be received from the channel
	fmt.Println(msg)
}
