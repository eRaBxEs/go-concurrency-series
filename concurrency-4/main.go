package main

import (
	"fmt"
	"sync"
)

func main() {
	// use a synchronisation primitive; declare variable of type mutex
	var m sync.Mutex

	ch := make(chan int) // create a channel

	defer close(ch) // to ensure that channel is closed

	go goRoutine1(ch, &m)
	ch <- 1 // Note that a value is sent to the channel after the first goRoutine1 is spawned.

}

func goRoutine1(ch chan int, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	for i := range ch {
		fmt.Println("Value received from channel in goRoutine1:", i)
	}
}

func goRoutine2(ch chan int, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	for i := range ch {
		fmt.Println("Value received from channel in goRoutine2:", i)
	}
}
