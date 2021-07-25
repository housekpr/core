package gpioctrl

import (
	"log"

	"github.com/warthog618/gpiod"
)

// global vars, todo: transform to struct once working...
var c *gpiod.Chip = nil        // GPIO chip
var lPumpOut *gpiod.Line = nil // GPIO line for pump control
// var lDistTrig *gpiod.Line = nil // GPIO line for distance sensor triger (RX)
// var lDistEcho *gpiod.Line = nil // GPIO line for distance sensor echo (TX)

// Prepare all GPIO chips & pins that may be used
func PrepareGPIO() {
	// error var. Creating explicitly, because chip and line already exist.
	var err error = nil

	if c == nil {
		// todo: move the name to config.
		chip := "gpiochip0"
		// Get chip0
		c, err = gpiod.NewChip(chip)
		if err == nil {
			log.Printf("Successfully opened GPIO character device '%s'\n", chip)
		} else {
			panic(err)
		}
	}

	// if lDistTrig == nil {
	// 	// todo: move pin selection to config.
	// 	distTrigPin := 15
	// 	// Request the pin as
	// 	lDistTrig, err = c.RequestLine(distTrigPin, gpiod.AsOutput())
	// 	if err == nil {
	// 		log.Printf("Successfully requested control for line 'GPIO%d'\n", distTrigPin)
	// 	} else {
	// 		panic(err)
	// 	}
	// }

	// if lDistEcho == nil {
	// 	// Request GPIO4 (PIN7) as output
	// 	dstEcho := 15
	// 	lPumpOut, err = c.RequestLine(dstEcho, gpiod.AsOutput())
	// 	if err == nil {
	// 		log.Printf("Successfully requested control for line 'GPIO%d'\n", dstEcho)
	// 	} else {
	// 		panic(err)
	// 	}
	// }
}

func ReleaseGPIO() {
	err := lPumpOut.Close()
	if err == nil {
		lPumpOut = nil
		log.Println("Successfully released control for line 'GPIO4'")
	} else {
		panic(err)
	}

	err = c.Close()
	if err == nil {
		c = nil
		log.Println("Successfully released control for chip 'gpiochip0'")
	} else {
		panic(err)
	}

}

// Get the reference to the chip
func GetChip() *gpiod.Chip {
	PrepareGPIO()
	return c
}

// Make sure the pump line is ready, before using it.
func preparePump() {
	PrepareGPIO()
	// error var. Creating explicitly, because chip and line already exist.
	var err error = nil

	if lPumpOut == nil {
		// todo: move pin selection to config.
		pumpPin := 4
		// Request the pin as output.
		lPumpOut, err = c.RequestLine(pumpPin, gpiod.AsOutput())
		if err == nil {
			log.Printf("Successfully requested control for line 'GPIO%d'\n", pumpPin)
		} else {
			panic(err)
		}
	}
}

// Get the reference to the line of the pump
func GetPumpLine() *gpiod.Line {
	preparePump()
	return lPumpOut
}

// Get the reference to the line of the pump
// func GetDistLines() (lDistTrig *gpiod.Line, lDistEcho *gpiod.Line) {
// 	PrepareGPIO()
// 	return lDistTrig, lDistEcho
// }
