package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const numberOfPizza = 10

var pizzasMade, pizzaFailed, total int

type producer struct {
	data chan pizzaOrder
	quit chan chan error
}

type pizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *producer) close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNumber int) *pizzaOrder {
	pizzaNumber++
	if pizzaNumber <= numberOfPizza {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d", pizzaNumber)
		rnd := rand.Intn(12) + 1
		msg := ""
		success := false
		if rnd < 5 {
			pizzaFailed++
		} else {
			pizzasMade++
		}
		total++
		fmt.Printf("Making pizza #%d, it wil take %d seconds..\n", pizzaNumber, delay)
		// delay for a bit

		time.Sleep(time.Duration(delay) * time.Second)
		if rnd < 2 {
			msg = fmt.Sprintf(" *** we run  out of ingredients for pizza #%d", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf(" *** cook quit while making pizza #%d", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf(" pizza #%d is ready", pizzaNumber)
		}
		p := pizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}
		return &p
	}
	return &pizzaOrder{pizzaNumber: pizzaNumber}
}

func pizzeria(pizzaMaker *producer) {
	// keep track which pizza we are making
	var i = 0
	// run forever or until we receive quit notification
	// try to make pizza
	for {
		currentPizza := makePizza(i)
		// try to make a pizza
		// decision if pizza was failed success or declined
		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit:
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	// get random number
	rand.Seed(time.Now().UnixNano())

	// print out message
	color.Cyan("The Pizzeria is open for business!")
	color.Cyan("----------------------------------")

	// create producer
	pizzaJob := &producer{
		data: make(chan pizzaOrder),
		quit: make(chan chan error),
	}

	// run the producer in the background
	go pizzeria(pizzaJob)

	// create consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizza {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("Customer is really mad")
			}
		} else {
			color.Cyan("Done making pizzas")
			err := pizzaJob.close()
			if err != nil {
				color.Red("Error closing channel", err)
			}
		}
	}
	color.Cyan("------------------------")
	color.Cyan("Done for today\n")
	color.Cyan("We've made %d pizzas, but failed to make %d, with %d attempts in total", pizzasMade, pizzaFailed, total)

}
