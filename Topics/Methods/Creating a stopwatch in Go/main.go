package main

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start time.Time
	end   time.Time
}

// Create the Start(), End() and TimeElapsed() methods below:
func (s *Stopwatch) Start() {
	s.start = time.Now()
}

func (s *Stopwatch) End() {
	s.end = time.Now()
}

func (s *Stopwatch) TimeElapsed() time.Duration {
	return s.end.Sub(s.start)
}

// DO NOT delete or modify the contents of the main() function!
func main() {
	var s Stopwatch
	var seconds int
	fmt.Scanln(&seconds)

	s.Start()
	time.Sleep(time.Duration(seconds) * time.Second)
	s.End()
	fmt.Printf("Elapsed time: %.3vs", s.TimeElapsed())
}
