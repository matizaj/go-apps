package main

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

// list of philosophers
var philosophers = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

var hunger = 3 // how many times a person eat
var eatTime = time.Second * 1
var thinkTime = time.Second * 3
var sleepTime = time.Second * 1

func main() {
	// hello message
	fmt.Println("Dining Philosopher problem")
	fmt.Println("--------------------------")
	fmt.Println("The table is empty")

	// start a meal
	dine()

	// end message
	fmt.Println("The table is empty")
}

func dine() {
	//
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	forks := make(map[int]*sync.Mutex)
	for i := 0; i < len(philosophers); i++ {
		forks[i] = &sync.Mutex{}
	}

	// start a meal
	for i := 0; i < len(philosophers); i++ {
		// fire of goorutine for current philisophers
		go diningProblem(philosophers[i], wg, forks, seated)
	}
	// Wait for the philosophers to finish. This blocks until the wait group is 0.
	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()
	// seat the philosophers at the table
	fmt.Printf("%s seated at the table\n", philosopher.name)
	seated.Done()
	seated.Wait()

	// eat 3 times
	for i := hunger; i > 0; i-- {
		// get lock for both forks
		if philosopher.leftFork > philosopher.rightFork {
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork\n", philosopher.name)
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork\n", philosopher.name)
		} else {
			forks[philosopher.leftFork].Lock()
			fmt.Printf("%s takes the left fork\n", philosopher.name)
			forks[philosopher.rightFork].Lock()
			fmt.Printf("%s takes the right fork\n", philosopher.name)
		}

		fmt.Printf("%s takes both forks and eat\n", philosopher.name)
		time.Sleep(eatTime)

		fmt.Printf("%s is thinking\n", philosopher.name)
		time.Sleep(thinkTime)

		forks[philosopher.leftFork].Unlock()
		forks[philosopher.rightFork].Unlock()
		fmt.Printf("%s put forks down\n", philosopher.name)
	}
	fmt.Printf("%s is satisfied\n", philosopher.name)
	fmt.Printf("%s left table\n", philosopher.name)
}
