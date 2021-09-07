package main

import "fmt"

var forks [5]fork

func philosopher(leftFork fork, rightFork fork) {
	for {
		// - is leftFork in use? AND is rightFork in use?
		// - if yes, grab them both and eat
		// - if no, leave them and think

		// ...

		// - put down both forks
	}
}

func main() {
	fmt.Println("Hello world")

	for i := range forks {
		forks[i] = fork{
			input:  make(chan int, 5),
			output: make(chan int, 5),
		}
	}

}
