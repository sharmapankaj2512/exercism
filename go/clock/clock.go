package clock

import (
	"fmt"
)

const testVersion = 4

type Clock struct {
	hh int
	mm int
}

func New(hour, minute int) Clock {
	mm, hh := MM(minute)
	return Clock{HH(hour, hh), mm}
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.hh, clock.mm)
}

func (clock Clock) Add(minutes int) Clock {
	return New(clock.hh, clock.mm + minutes)
}

func HH(hour, hh int) int {
	h := (hour + hh) % 24
	if h < 0 {
		return 24 + h
	}
	return h
}

func MM(minute int) (mm, hh int) {
	mm = minute % 60
	hh = minute / 60
	if mm < 0 {
		return 60 + mm, hh - 1
	}
	return mm, hh
}
