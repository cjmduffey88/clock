package clock

import "time"

type Stepper interface {
	Step()
}

type Clock struct {
	frequency uint
	steppers  []Stepper
	ticker    *time.Ticker
}

func NewClock(frequency uint) *Clock {
	clock := new(Clock)
	clock.frequency = frequency
	clock.ticker = time.NewTicker(time.Second / time.Duration(frequency))
	return clock
}

func (clock *Clock) AddStepper(stepper Stepper) {
	clock.steppers = append(clock.steppers, stepper)
}

func (clock *Clock) Run() {
	defer clock.ticker.Stop()
	for range clock.ticker.C {
		for _, stepper := range clock.steppers {
			stepper.Step()
		}
	}
}
