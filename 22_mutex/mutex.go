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
	// defer wg.Done()
	// defer p.mu.Unlock()
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1
	// p.mu.Unlock() //recommended to put this unlock inside defer so that for more operations it executes at the end of function execution
}

// mutex: locking of resource, we do it usually with concurrent with go routines. It solves race conditions

func main() {
	var wg sync.WaitGroup

	myPost := post{views: 0}
	for range 1000 {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.views)
}
