package clock

import (
	"fmt"
)

const testVersion = 4

type Clock int

const totalMinutes = 24 * 60

func New(hours, minutes int) Clock {
	return Clock(positiveMod(hours * 60 + minutes))
}

func positiveMod(minutes int) int {
	return (minutes % totalMinutes + totalMinutes) % totalMinutes
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock / 60, clock % 60)
}

func (clock Clock) Add(minutes int) Clock {
	return Clock(positiveMod(int(clock) + minutes))
}
