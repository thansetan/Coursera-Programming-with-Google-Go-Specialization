package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Implement the dining philosopher's problem with the following constraints/modifications:

 1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
 2. Each philosopher should eat only 3 times.
 3. The philosophers pick up the chopsticks in any order,
 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
 5. The host allows no more than 2 philosophers to eat concurrently.
 6. Each philosopher is numbered, 1 through 5.
 7. When a philosopher starts eating (after it has obtained necessary locks) prints "starting to eat <number>" on a line by itself,
    where <number> is the number of the philosopher.
 8. When a philosopher finishes eating (before it has released its locks) prints "finishing to eat <number>" on a line by itself,
    where <number> is the number of the philosopher.
*/

type Philosopher struct {
	id, eatingCount               int
	leftChopstick, rightChopstick *sync.Mutex
}

func (p *Philosopher) eat(hostCh chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for p.eatingCount < 3 {
		hostCh <- struct{}{}

		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("Philosopher %d starting to eat\n", p.id)

		time.Sleep(5 * time.Millisecond)
		p.eatingCount++

		fmt.Printf("Philosopher %d finishing to eat\n", p.id)
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

	}
}

func main() {
	var wg sync.WaitGroup
	hostCh := make(chan struct{}, 2)
	go host(hostCh)

	chopsticks := make([]sync.Mutex, 5)

	philosophers := make([]Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = Philosopher{
			id:             i + 1,
			leftChopstick:  &chopsticks[i],
			rightChopstick: &chopsticks[(i+1)%5],
		}
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go philosophers[i].eat(hostCh, &wg)
	}

	wg.Wait()
}

func host(hostCh chan struct{}) {
	for {
		if len(hostCh) == 2 {
			<-hostCh
			<-hostCh
		}
	}
}
