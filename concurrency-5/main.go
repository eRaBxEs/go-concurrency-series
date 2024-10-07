package main

import "fmt"

func main() {
	// create a buffered channel of size 3
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
		case charChannel <- char: // buffered channel is non blocking. They are asynchronus
		}
	}

	close(charChannel)
	// note that even after closing the channel above, one can still get the values in the channel
	for result := range charChannel {
		fmt.Println(result)
	}
}
