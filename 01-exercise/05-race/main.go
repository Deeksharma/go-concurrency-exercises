package main

import (
	"fmt"
	"math/rand"
	"time"
)

//TODO: identify the data race
// fix the issue.

func main() {
	start := time.Now()
	var ch chan int
	ch = make(chan int)
	defer close(ch)
	var t *time.Timer //main go routine access t
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- 1
	})
	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}
	//t.Reset(randomDuration()) // timeAfterGoroutine access t
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

//----------------------------------------------------
// (main goroutine) -> t <- (time.AfterFunc goroutine)
//----------------------------------------------------
// (working condition)
// main goroutine..
// t = time.AfterFunc()  // returns a timer..

// AfterFunc goroutine
// t.Reset()        // timer reset
//----------------------------------------------------
// (race condition- random duration is very small)
// AfterFunc goroutine
// t.Reset() // t = nil

// main goroutine..
// t = time.AfterFunc()
//----------------------------------------------------

/*
The variable t, is being accessed by two goroutines, the main goroutine and the time after func goroutines,

in a working conditions, the time after func is going to assign a value to variable t, and time after

func goroutine is going to use that variable to reset the timer.

But what happens if the random duration that is generated is very small, in that case, it is very much

possible that the time after func goroutine gets scheduled, before the value is assigned to the variable t.

In that case, the time after func goroutine will be trying to reset the timer on nil value, which

can lead to program crashing.


*/

/*
fix??

We can confine the read and write to the variable t, to only one goroutine.

In this case, the main goroutine, so that the read and the write to the variable t, is synchronized.

And we can use a channel for communication between the goroutine created by the time after function and

the main goroutine.

So when the time after function wants a timer to be reset, it's going to send a message to the main

goroutine and the main goroutine is going to reset the timer.

*/

// Imp - confine the scope of variable to one goroutine
