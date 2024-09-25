package main

import (
	"fmt"
	"sync"
)

func main() {
	// use a synchronisation primitive; declare variable of type mutex
	var m sync.Mutex

	ch := make(chan int) // create a channel

}

func goRoutine1(ch chan int, m *sync.Mutex) {
	m.Lock()
	defer m.Unlock()
	for i := range ch {
		fmt.Println("Value received from channel in goRoutine1:", i)
	}
}
