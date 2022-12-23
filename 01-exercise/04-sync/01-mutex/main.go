//package main
//
//import (
//	"fmt"
//	"runtime"
//	"sync"
//)
//
//func main() {
//
//	runtime.GOMAXPROCS(4)
//
//	var balance int // balance variable is a shared variable between the goroutines, the addition and subtraction operation are not atomic, thats why we are getting wrong result
//	var wg sync.WaitGroup
//
//	deposit := func(amount int) {
//		balance += amount
//	}
//
//	withdrawal := func(amount int) {
//		balance -= amount
//	}
//
//	// make 100 deposits of $1
//	// and 100 withdrawal of $1 concurrently.
//	// run the program and check result.
//
//	// TODO: fix the issue for consistent output.
//
//	wg.Add(100)
//	for i := 0; i < 100; i++ {
//		go func() {
//			defer wg.Done()
//			deposit(1)
//		}()
//	}
//
//	wg.Add(100)
//	for i := 0; i < 100; i++ {
//		go func() {
//			defer wg.Done()
//			withdrawal(1)
//		}()
//	}
//
//	wg.Wait()
//	fmt.Println(balance)
//}

package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
It involves retrieving the value of the balance from memory, doing the operation addition

or subtraction, and then storing the value of the balance to the memory.

# So if the goroutines are running concurrently in interleaved manner, preempting each other, then data race

is possible and data can get corrupted.

So we need to use mutex to guard the access to the shared variable from multiple goroutines.
*/
func main() {

	runtime.GOMAXPROCS(4)

	var balance int // balance variable is a shared variable between the goroutines, the addition and subtraction(3 steps) operation are not atomic, thats why we are getting wrong result
	var wg sync.WaitGroup
	var mu sync.Mutex

	deposit := func(amount int) {
		mu.Lock()
		balance += amount // critical section of code - use mutex to guard te shared variable
		mu.Unlock()
	}

	withdrawal := func(amount int) {
		mu.Lock()
		defer mu.Unlock()
		balance -= amount
	}

	// make 100 deposits of $1
	// and 100 withdrawal of $1 concurrently.
	// run the program and check result.

	// TODO: fix the issue for consistent output.

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}

	wg.Wait()
	fmt.Println(balance)
}
