package main

import (
	"fmt"
	"math/rand"
)

type SlotMachine struct {
	Slots [3]int
}

const nineUnits = 9

// Create the Play() and CheckWin() methods below:
func (s *SlotMachine) Play() {
	for i := range s.Slots {
		s.Slots[i] = rand.Intn(nineUnits)
	}
}

func (s *SlotMachine) CheckWin() string {
	if s.Slots[0] == s.Slots[1] && s.Slots[0] == s.Slots[2] {
		return "Jackpot!"
	}
	return "You lost!"
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var money int64
	fmt.Scanln(&money)
	rand.Seed(money)

	var sm SlotMachine
	sm.Play()
	fmt.Println(sm.Slots)
	fmt.Println(sm.CheckWin())
}
