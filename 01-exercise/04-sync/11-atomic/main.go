package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4)

	var counter uint64 // this counter is accessed by multiple goroutines parellaly, so the data can be corrupted
	// So there will be a data race and the data can get corrupted.
	var wg sync.WaitGroup

	// TODO: implement concurrency safe counter

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("counter: ", counter)
}
