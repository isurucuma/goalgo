package main

// this is a implementation of the dining philosophers problem but with a slight modification of using a host
// here the constraint is that only 2 philosophers can eat at a time
// so the host is created with a buffered channel of size 2
import (
	"fmt"
	"sync"
	"time"
)

// Chopstick is simply a mutex, all the methods of mutex can be used on it
type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	host           *Host
	timesToEat     int
}

// here a buffered channel is used to define the host
// maximum upto 2 philosophers can add to the channel
// others have to wait until atleast one slot in the buffered channel is free
type Host struct {
	permission chan int
}

func (host *Host) allow() {
	host.permission <- 1
}

func (host *Host) release() {
	<-host.permission
}

func (p *Philosopher) eat() {
	for i := 0; i < p.timesToEat; i++ {
		p.host.allow()
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("starting to eat %d\n", p.number)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("finishing eating %d\n", p.number)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()
		p.host.release()
	}
}

func main() {
	chopsticks := make([]*Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(Chopstick)
	}

	// only a single host is created with a buffered channel of size 2
	host := &Host{
		permission: make(chan int, 2),
	}

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{
			number:         i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%5],
			host:           host, // same host os passed to all philosophers
			timesToEat:     3,
		}
	}

	var wg sync.WaitGroup
	wg.Add(5)
	for _, p := range philosophers {
		go func(p *Philosopher) {
			defer wg.Done()
			p.eat()
		}(p)
	}
	wg.Wait()
}
