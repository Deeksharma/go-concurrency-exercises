package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	m := sync.Mutex{}
	c := sync.NewCond(&m)
	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		c.L.Lock()
		// this all needs to be inside lock - like one operation
		for len(sharedRsc) == 0 {
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	// writes changes to sharedRsc
	c.L.Lock() // since updating the shared resource
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()

	wg.Wait()
}
