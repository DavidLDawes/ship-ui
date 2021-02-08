package main

import (
	"fyne.io/fyne/widget"
)

var tlSelect *widget.Select

func techLevelChanged(tlSelected string) {
	if len(tlSelected) == 1 {
		StarShip.tl = tlSelected
		tlOffset := tlSelected[0] - 70
		// Odd one, skips I so we have:
		// F = 0
		// G = 1
		// H = 2
		// I
		// J = 3
		// K = 4
		// L = 5
		if tlOffset < 7 {
			StarShip.tl = tlSelected
			StarShip.tlOffset = int(tlOffset)
			if StarShip.tlOffset > 2 {
				StarShip.tlOffset--
			}
			if StarShip.tlOffset < 1 {
				// TL-F (offset 0) only goes to J-6, M-6 & P-6
				if StarShip.jump > 6 {
					StarShip.jump = 6
					jumpSelect.SetSelected("6")
				}
				if StarShip.maneuver > 6 {
					StarShip.maneuver = 6
					maneuverSelect.SetSelected("6")
				}
				if StarShip.power > 6 {
					StarShip.power = 6
					powerSelect.SetSelected("6")
				}
			}

		}
		buildJump()
		buildFuel()
		buildManeuver()
		buildPower()
		buildBridge()
		buildHardPoints()
		buildTotal()
		adjustSlider()
	}
}
