package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	//TODO: fix the issue.

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) { // if we dont pass i the go routine will tak the value that is present in the heap i,e = 4, if we want the current value of i tobe printed we need to pass it as a variable
			defer wg.Done()
			fmt.Println(i)
		}(i)

	}
	wg.Wait() // adding wait group after the wg.Add will simply act like running the code swequentially
}
