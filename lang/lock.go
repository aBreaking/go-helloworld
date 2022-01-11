package main

import "sync"

func main() {

}

type Info struct {
	mu  sync.Mutex
	Str string
}

func Update(info *Info) {
	info.mu.Lock()
	// critical section:
	info.Str = "abc" // new value
	// end critical section
	info.mu.Unlock()
}
