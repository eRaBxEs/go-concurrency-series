package main

import (
	"fmt"
)

func main() {
	myChannel := make(chan string)
	anotherChannel := make(chan string)

	go func() {
		myChannel <- "data"
	}()

	go func() {
		anotherChannel <- "foo"
	}()

	// using a select statement which will be used to select or get a message from a channel which is ready at that moment
	// if the two messages are ready at the same time, it picks one randomly
	// note that the select is also going to block till it receives a message from either of the channel

	select {
	case messageFromMychannel := <-anotherChannel:
		fmt.Println(messageFromMychannel)

	case messageFromAnotherChannel := <-anotherChannel:
		fmt.Println(messageFromAnotherChannel)
	}
}
