package sampler

import (
	"fmt"
	"time"
)

type Sampler struct {
	duration time.Duration
	counter  int

	perFrames int
	msg       string
}

func New(perFrames int, msg string) *Sampler {
	return &Sampler{perFrames: perFrames, msg: msg}
}

func (s *Sampler) Sample(fu func()) {
	t := time.Now()
	fu()
	s.duration += time.Since(t)
	s.counter++
	if s.counter%s.perFrames == 0 {
		d := s.duration
		fmt.Printf("%s: %v\n", s.msg, d/time.Duration(s.counter))
		s.counter = 0
		s.duration = 0
	}
}
