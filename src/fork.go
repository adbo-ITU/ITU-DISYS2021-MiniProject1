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
