package day23

import (
	"fmt"
	"strings"
	"time"

	"github.com/domszyn/adventofcode/2019/toolbox"
)

type Packet struct {
	Destination int
	X           int
	Y           int
}

func solvePart1() (int, int) {
	masterProgram := toolbox.LoadProgram(strings.TrimSpace(Input))

	packetQueue := make(chan Packet, 256)
	inputs := make([]chan int, 50)
	outputs := make([]chan int, 50)
	done := make(chan bool, 50)

	for i := range 50 {
		// Use a buffer of 2 to be safe, in case a computer outputs
		// an address and a value before the manager can read them.
		inputs[i] = make(chan int, 256)
		outputs[i] = make(chan int, 256)
	}

	programs := make([]*toolbox.Program, 50)
	for i := range 50 {
		p := masterProgram.Copy()
		programs[i] = &p
	}

	for i := range 50 {
		go programs[i].IntCode(inputs[i], outputs[i], done)
		// Tell the computer its network address
		inputs[i] <- i
	}

	for i := range 50 {
		go func(computerID int) {
			for {
				select {
				case dest := <-outputs[computerID]:
					x := <-outputs[computerID]
					y := <-outputs[computerID]
					packetQueue <- Packet{Destination: dest, X: x, Y: y}
				case <-done:
					return
				}
			}
		}(i)
	}

	var part1 int
	var computersDone int
	var natPacket Packet
	var lastDelivered int
routerLoop:
	for {
		select {
		case packet := <-packetQueue:
			fmt.Printf("%+v\n", packet)
			if packet.Destination == 255 {
				if part1 == 0 {
					part1 = packet.Y
				}
				natPacket = packet
			} else if packet.Destination < 50 {
				inputs[packet.Destination] <- packet.X
				inputs[packet.Destination] <- packet.Y
			}
		case <-time.After(1000 * time.Millisecond):
			if natPacket.X == 0 && natPacket.Destination == 0 {
				continue
			}

			if natPacket.Y == lastDelivered && lastDelivered != 0 {
				return part1, natPacket.Y
			}

			inputs[0] <- natPacket.X
			inputs[0] <- natPacket.Y
			lastDelivered = natPacket.Y
		case <-done:
			computersDone++
			if computersDone == 50 {
				break routerLoop
			}
		}
	}

	return part1, -1
}

func GetAnswers() (int, int) {
	return solvePart1()
}
