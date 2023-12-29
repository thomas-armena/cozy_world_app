package clock

type FakeClock struct {
	Clock
	currentTimeUsec int64
}

func NewFakeClock(startTimeUsec int64) *FakeClock {
	return &FakeClock{
		currentTimeUsec: startTimeUsec,
	}
}

func (f *FakeClock) SetTime(newTimeUsec int64) {
	f.currentTimeUsec = newTimeUsec
}

func (f *FakeClock) AdvanceUsec(usec int64) {
	f.currentTimeUsec += usec
}

func (f *FakeClock) AdvanceMs(msec int64) {
	f.currentTimeUsec += msec * 1000
}

func (f *FakeClock) AdvanceSec(sec int64) {
	f.currentTimeUsec += sec * 1000000
}

func (f *FakeClock) CurrentTime() int64 {
	return f.currentTimeUsec
}
