package main

import (
	"fmt"
	"sync"
)

/*
Write two goroutines which have a race condition when executed concurrently.
Explain what the race condition is and how it can occur.
*/

/*
Race condition is a condition where multiple concurrent operations (in this case goroutine) access
and manipulate a shared data without proper synchronization. In the code below, both goroutine access
the same variable (x). The first goroutine will increase the value of x 100000 times, and the second one
will decrease the value of x 100000 times. If you roun this code multiple times, the value of x will be
different by the end of the execution.
*/
func main() {
	var x int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		var y int
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mu.Lock()
			x++
			mu.Unlock()
			y++
		}
		fmt.Printf("This goroutine increase the value of x %d times\n", y)
	}()

	go func() {
		var y int
		defer wg.Done()
		for i := 0; i < 100000; i++ {
			mu.Lock()
			x--
			mu.Unlock()
			y++
		}
		fmt.Printf("This goroutine decrease the value of x %d times\n", y)
	}()

	wg.Wait()
	fmt.Printf("value of x : %d\n", x)
}
