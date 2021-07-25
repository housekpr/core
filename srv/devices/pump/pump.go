package pump

import "github.com/housekpr/core/srv/devices/gpioctrl"

func ReadState() int {
	l := gpioctrl.GetPumpLine() // Get the pump gpio line reference

	r, _ := l.Value() // Read the line value

	return r // Return the result
}

func Start() {
	l := gpioctrl.GetPumpLine() // Get the pump gpio line reference

	l.SetValue(1) // Change the value to true (1) to start the pump
}

func Stop() {
	// Get Pump Line
	l := gpioctrl.GetPumpLine()

	l.SetValue(0) // Change the value to false (0) to stop the pump
}
