package fork

import "fmt"

type fork struct {
	input  chan<- bool // <- action: do i toggle the picked up status?
	output <-chan ForkAnswer // <- info: times picked up, is it currently picked up
}

// Eternal loop - locks fork until its updated - max 1 user at a time because write/read blocks
for {
	skriv til output //Sender fork status til EN philosopher
	læs fra input //Modtager eventuel statusændring fra den samme philosopher
	opdater værdier
}

type ForkAnswer struct {
	question  string
	pickedUp  bool
	timesUsed uint64
}

func Fork(forkInfo fork) {
	fmt.Println("Hello World")
}
