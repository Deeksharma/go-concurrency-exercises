package main

import (
	"fmt"
	"sync"
)

//TODO: run the program and check that variable i
// was pinned for access from goroutine even after
// enclosing function returns.

func main() {
	var wg sync.WaitGroup

	incr := func(wg *sync.WaitGroup) {
		var i int // local variable i
		wg.Add(1)
		go func() {
			defer wg.Done()
			i++
			fmt.Printf("value of i: %v\n", i)
		}()
		fmt.Println("return from function")
		return // here the function is retruning still the reference to the variable i is being held
	}

	incr(&wg)
	wg.Wait()
	fmt.Println("done..")
}
