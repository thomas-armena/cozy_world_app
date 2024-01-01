package clock

import "time"

type Clock interface {
	CurrentTime() int64
}

type SystemClock struct{}

func NewSystemClock() *SystemClock {
	return &SystemClock{}
}

func (s *SystemClock) CurrentTime() int64 {
	return time.Now().UnixMicro()
}
