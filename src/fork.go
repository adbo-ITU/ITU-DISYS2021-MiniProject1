package main

type fork struct {
	isPickedUp bool
	numPickUps uint64
	input      chan int        // <- action: do i toggle the picked up status?
	output     chan ForkStatus // <- info: times picked up, is it currently picked up
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
