package main

import (
	"strconv"

	"fyne.io/fyne/widget"
)

var tlSelect *widget.Select
var tonsSelect *widget.Select

func tlChanged(tlSelected string) {
	if len(tlSelected) == 1 {
		tlOffset := tlSelected[0] - 46
		// Odd one, skips I so we have:
		// F = 0
		// G = 1
		// H = 2
		// J = 3
		// K = 4
		if tlOffset < 6 {
			if tlOffset > 4 {
				tlOffset = 4
			}
			StarShip.tl = tlSelected
			StarShip.tlOffset = int(tlOffset)
		}
	}

}

func techLevelChanged(value string) {
	tech, err := strconv.Atoi(value)
	if err == nil {
		if tech < 6 {
			if tech > 4 {
				tech = 4
			}
			if tech < 1 {
				// TL-F (offset 0) only goes to J-6, M-6 & P-6
				if StarShip.jump > 6 {
					StarShip.jump = 6
				}
				if StarShip.maneuver > 6 {
					StarShip.maneuver = 6
				}
				if StarShip.power > 6 {
					StarShip.power = 6
				}
			}
			StarShip.tl = value
			StarShip.tlOffset = int(tech)
			buildJump()
			buildFuel()
			buildManeuver()
			buildPower()
			buildBridge()
			buildHardPoints()
			BuildTotal()
		}
	}
}
