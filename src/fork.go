package main

type fork struct {
	isPickedUp bool
	numPickUps uint64
	input      chan int
	output     chan ForkStatus
}

type ForkStatus struct {
	isPickedUp bool
	numPickUps uint64
}

func (f fork) forkLife() {
	for {
		// Design decision: interacting with forks requires to read their status
		// and later telling them what the single philosopher will do with them

		f.output <- ForkStatus{isPickedUp: f.isPickedUp, numPickUps: f.numPickUps}
		action := <-f.input

		if action == pickUp {
			f.isPickedUp = true
			f.numPickUps++
		} else if action == putDown {
			f.isPickedUp = false
		}
	}
}

func NewFork() fork {
	inputChannel := make(chan int)
	outputChannel := make(chan ForkStatus)
	return fork{isPickedUp: false, numPickUps: 0, input: inputChannel, output: outputChannel}
}
