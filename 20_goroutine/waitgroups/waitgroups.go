package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println("doing task", id)
}

func main() {
	// creating wait group
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		// go task(i, &wg)

		// for inline function use
		func() {
			defer wg.Done()
			fmt.Println("doing task", i)
		}()
	}

	wg.Wait()
}
