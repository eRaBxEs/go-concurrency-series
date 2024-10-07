package main

import "fmt"

func main() {
	// create a buffered channel of size 3
	charChannel := make(chan string, 3)
	chars := []string{"a", "b", "c"}

	for _, char := range chars {
		select {
		case charChannel <- char:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}
