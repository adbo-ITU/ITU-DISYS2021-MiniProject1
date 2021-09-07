package main

import (
	"fmt"
	"time"
)

var forks [5]fork
var philosophers [5]philosopher

func main() {
	// initialize the forks
	for i, _ := range forks {
		inputBuffer := make(chan int)
		outputBuffer := make(chan ForkStatus)
		forkBuffer := fork{isPickedUp: false, numPickUps: 0, input: inputBuffer, output: outputBuffer}
		forks[i] = forkBuffer
		go forkBuffer.forkLife()
	}

	// initialize the philosophers
	for i, _ := range philosophers {
		philosopherBuffer := philosopher{isEating: false, numEats: 0, numThinks: 0}
		philosophers[i] = philosopherBuffer
		go philosophers[i].philosopherLife(forks[i], forks[(i+1)%len(philosophers)])
	}

	// prevent the main thread from exiting
	for {
		// clear terminal window
		fmt.Print("\033[H\033[2J")

		// print philosopher info
		for i, phil := range philosophers {
			fmt.Printf("[STATUS]: Philosopher %d: has eaten %d times, has been thinking %d times\n",
				i,
				phil.numEats,
				phil.numThinks)
		}

		time.Sleep(500 * time.Millisecond)
	}
}
