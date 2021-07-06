package pump

import (
	"github.com/warthog618/gpiod"
)

var c *gpiod.Chip = nil // GPIO chip global var
var l *gpiod.Line = nil // GPIO pump control line global var

func prepareGPIO() {
	var err error = nil // var to check for errors.

	if c == nil {
		// Get chip0
		c, err = gpiod.NewChip("gpiochip0")
		if err != nil {
			panic(err)
		}
	}

	if l == nil {
		// Request GPIO4 (PIN7) as output
		l, err = c.RequestLine(4, gpiod.AsOutput())
		if err != nil {
			panic(err)
		}
	}
}

func releaseGPIO() {
	// TODO - figure out how to close properly
	// l.Close()
	// c.Close()
	// l = nil
	// c = nil
}

func ReadState() int {
	prepareGPIO()

	// Read Value
	r, _ := l.Value()

	releaseGPIO()
	// return the GPIO Line Value
	// inverting because of the relay
	if r == 1 {
		return 0
	} else {
		return 1
	}

}

func Start() {
	prepareGPIO()

	// Set the GPIO Line Value to 0 (ON)
	// inverting because of the relay
	l.SetValue(0)

	releaseGPIO()
}

func Stop() {
	prepareGPIO()

	// Set the GPIO Line Value to 1 (OFF)
	// inverting because of the relay
	l.SetValue(1)

	releaseGPIO()
}
