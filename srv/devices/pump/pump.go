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
	return r

}

func Start() {
	prepareGPIO()

	l.SetValue(1)

	releaseGPIO()
}

func Stop() {
	prepareGPIO()

	l.SetValue(0)

	releaseGPIO()
}
