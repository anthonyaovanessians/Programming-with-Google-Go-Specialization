package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type chopstick struct {
	sync.Mutex
}
type philo struct {
	number            int
	left_cs, right_cs *chopstick
}

func main() {

	chopsticks := make([]*chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = new(chopstick)
	}

	philos := make([]philo, 5)
	for i := 0; i < 5; i++ {
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		var a int
		var b int
		for {
			a = r1.Intn(5)
			b = r1.Intn(5)
			if a == b {
				continue
			} else {
				break
			}
		}
		philos[i] = philo{i + 1, chopsticks[a], chopsticks[b]}
	}

	for i := 0; i < 5; i++ {
		var wg sync.WaitGroup
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		var a int

		for {
			a = r1.Intn(5)
			if i == a {
				continue
			} else {
				break
			}
		}
		wg.Add(1)
		go permission(philos[i], philos[a], &wg)
		wg.Wait()
	}
}

func permission(p1, p2 philo, wg *sync.WaitGroup) {
	var new_wg sync.WaitGroup

	new_wg.Add(2)
	go p1.eat(&new_wg)
	go p2.eat(&new_wg)
	new_wg.Wait()

	wg.Done()

}

func (p philo) eat(wg *sync.WaitGroup) {

	fmt.Printf("May I, philosopher %v, eat?\n", p.number)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	input := r1.Intn(2)

	if input == 1 {
		fmt.Printf("Philosopher %v, you may eat.\n", p.number)

		for i := 0; i < 3; i++ {

			p.left_cs.Lock()
			p.right_cs.Lock()

			fmt.Printf("starting to eat %v\n", p.number)
			fmt.Printf("finishing eating %v\n", p.number)

			p.left_cs.Unlock()
			p.right_cs.Unlock()

		}

	} else {
		fmt.Printf("Sorry philosopher %v, you may not eat.\n", p.number)
	}
	fmt.Println()
	wg.Done()

}
