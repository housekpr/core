package distance

import (
	"log"
	"syscall"
	"time"

	"github.com/housekpr/core/srv/devices/gpioctrl"
	"github.com/warthog618/gpiod"
)

var RisingTime time.Time
var dist float64

var lDistTrig *gpiod.Line = nil // GPIO line for distance sensor triger (RX)
var lDistEcho *gpiod.Line = nil // GPIO line for distance sensor echo (TX)

// Watches GPIO15 (Raspberry Pi J8-10) and reports when it changes state.
func eventHandler(evt gpiod.LineEvent) {
	t := time.Now()

	if evt.Type == gpiod.LineEventFallingEdge {
		diff := t.Sub(RisingTime)
		dist = float64(diff.Microseconds()) / 58.0
		// log.Printf("[distance.go] LineEventFallingEdge received. diff: %v; dist: %v\n", diff, dist)
	} else {
		// edge := "rising"
		RisingTime = t
	}
	// log.Printf("event:%3d %-7s %s (%s)\n", evt.Offset, edge, t.Format(time.RFC3339Nano), evt.Timestamp)
}

// Gets the tank level
func GetTankLevel() float64 {
	// error var. Creating explicitly, because chip and line already exist.
	var err error = nil

	// Get the chip
	c := gpioctrl.GetChip()

	// Prepare the echo line. todo: move to function
	if lDistEcho == nil {
		echoPin := 14 // GPIO pin number
		lDistEcho, err = c.RequestLine(echoPin,
			gpiod.WithPullUp,
			gpiod.WithBothEdges,
			gpiod.WithEventHandler(eventHandler))

		if err == nil {
			log.Printf("[distance.go] Successfully requested control for line 'GPIO%d'", echoPin)
		} else {
			if err == syscall.Errno(22) {
				log.Println("[distance.go] Note that the WithPullUp option requires kernel V5.5 or later - check your kernel version.")
			} else {
				log.Printf("[distance.go] RequestLine returned error: %s\n", err)
			}
			panic(err)
		}
		// defer lDistEcho.Close()
	}

	// Prepare the trigger line. todo: move to function
	if lDistTrig == nil {
		triggerPin := 15 // GPIO pin number
		lDistTrig, err = c.RequestLine(triggerPin, gpiod.AsOutput(0))
		if err == nil {
			log.Printf("[distance.go] Successfully requested control for line 'GPIO%d'", triggerPin)
		} else {
			panic(err)
		}
	}
	// defer func() {
	// 	lDistTrig.Reconfigure(gpiod.AsInput)
	// 	lDistTrig.Close()
	// }()

	// log.Printf("Watching Pin %d...\n", offset)
	log.Println("[distance.go] Waiting 1sec in low level, to ensure stable measurement.")
	lDistTrig.SetValue(0)
	time.Sleep(1 * time.Second)

	//for {
	//log.Println("Measuring every 3sec...")
	log.Println("[distance.go] Initiating tank level measurement.")
	lDistTrig.SetValue(1)
	time.Sleep(20 * time.Microsecond)
	lDistTrig.SetValue(0)
	time.Sleep(1 * time.Second)
	log.Printf("[distance.go] Completed tank level measurement. Distance: %v", dist)
	return dist
	//}
}
