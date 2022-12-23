package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var x sync.Once

	load := func() {
		fmt.Println("Run only once initialization function")
	}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()

			//TODO: modify so that load function gets called only once.
			x.Do(load)
		}()
	}
	wg.Wait()
}
