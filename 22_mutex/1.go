package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()
	p.mu.Lock() //lock this before modification
	p.views++
	// p.mu.Unlock() //unlock after modification --> better to mive it to defer function so that it unlocks even in case of error etc
}

func main() {
	myPost := post{views: 0}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println("Post Views: ", myPost.views)
}
