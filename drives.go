package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

var detailJump widget.Label
var detailJumpFuel widget.Label
var detailManeuver widget.Label
var detailPower widget.Label

var jumpSelect *widget.Select
var maneuverSelect *widget.Select
var powerSelect *widget.Select

func jumpChanged(value string) {
	jump, err := strconv.Atoi(value)
	if err == nil {
		if StarShip.tlOffset < 1 {
			// TL-F (offset 0) only goes to J-6
			if jump > 6 {
				jump = 6
				jumpSelect.SetSelected("6")
			}
		}
		if jump > StarShip.power {
			jump = StarShip.power
			jumpSelect.SetSelected(fmt.Sprintf("%d", jump))
		}
		StarShip.jump = jump
	}
	StarShip.computer = computer[jump]
	buildJump()
	buildCrew()
	buildTotal()
}

func maneuverChanged(value string) {
	maneuver, err := strconv.Atoi(value)
	if err == nil {
		if StarShip.tlOffset < 1 {
			// TL-F (offset 0) only goes to M-6
			if maneuver > 6 {
				maneuver = 6
				maneuverSelect.SetSelected("6")
			}
		}
		if maneuver > StarShip.power {
			maneuver = StarShip.power
			maneuverSelect.SetSelected(fmt.Sprintf("%d", maneuver))
		}
		StarShip.maneuver = maneuver
	}
	buildManeuver()
	buildCrew()
	buildTotal()
}

func powerChanged(value string) {
	power, err := strconv.Atoi(value)
	if err == nil {
		if StarShip.tlOffset < 1 {
			// TL-F (offset 0) only goes to P-6
			if power > 6 {
				power = 6
				powerSelect.SetSelected("6")
			}
		}
		StarShip.power = power
		if power < StarShip.jump {
			StarShip.jump = power
			jumpSelect.SetSelected(fmt.Sprintf("%d", StarShip.jump))
			buildJump()
		}
		if power < StarShip.maneuver {
			StarShip.maneuver = power
			maneuverSelect.SetSelected(fmt.Sprintf("%d", StarShip.maneuver))
			buildManeuver()
		}
	}
	buildPower()
	buildCrew()
	buildTotal()
}

var jbase = [5]float32{
	2.0, 1.75, 1.5, 1.0, 0.66666,
}
var jinc = [5]float32{
	1.0, .8, .66666, 0.5, 0.33333,
}

func jumpRate() float32 {
	return (jbase[StarShip.tlOffset] + jinc[StarShip.tlOffset]*(float32(StarShip.jump)-1.0)) / 100.0
}

var jfbase = [5]float32{
	10.0, 8.0, 6.0, 5.0, 4.0,
}
var jfinc = [5]float32{
	10.0, 7.0, 6.0, 4.0, 2.5,
}

func jumpFuelRate() float32 {
	return (jfbase[StarShip.tlOffset] + jfinc[StarShip.tlOffset]*(float32(StarShip.jump)-1.0)) / 100.0
}

var mbase = [5]float32{
	2.0, 1.5, 1.25, 1.0, 0.66666,
}
var minc = [5]float32{
	3.0, 2.0, 1.66666, 1.5, 1.0,
}

func maneuverRate() float32 {
	return (mbase[StarShip.tlOffset] + minc[StarShip.tlOffset]*(float32(StarShip.maneuver)-1.0)) / 100.0
}

var pbase = [5]float32{
	1.0, 0.75, 0.6, 0.5, 0.33333,
}
var pinc = [5]float32{
	1.0, .7, 0.55, 0.4, 0.25,
}

func powerRate() float32 {
	return (pbase[StarShip.tlOffset] + pinc[StarShip.tlOffset]*(float32(StarShip.power)-1.0)) / 100.0
}

func buildJump() {
	StarShip.jumpTons = float32(StarShip.tons) * jumpRate()
	detailJump.SetText(fmt.Sprintf("Jump: %d, tons: %2.1f", StarShip.jump, StarShip.jumpTons))
	detailJump.Refresh()
	detailComputer.SetText(fmt.Sprintf("computer %d: %d tons", StarShip.jump, computer[StarShip.jump-1]))
	detailComputer.Refresh()
	setEngineers()
	refreshEngineeringCrew()
}
func buildManeuver() {
	StarShip.maneuverTons = float32(StarShip.tons) * maneuverRate()
	detailManeuver.SetText(fmt.Sprintf("Maneuver: %d, tons: %2.1f", StarShip.maneuver, StarShip.maneuverTons))
	detailManeuver.Refresh()
	setEngineers()
	refreshEngineeringCrew()
}
func buildPower() {
	StarShip.powerTons = float32(StarShip.tons) * powerRate()
	detailPower.SetText(fmt.Sprintf("Power: %d, tons: %2.1f", StarShip.power, StarShip.powerTons))
	detailPower.Refresh()
	setEngineers()
	refreshEngineeringCrew()
}
func buildFuel() {
	StarShip.jumpFuel = int(float32(StarShip.tons) * jumpFuelRate())
	detailJumpFuel.SetText(fmt.Sprintf("Jump fuel: %d", StarShip.jumpFuel))
	detailJumpFuel.Refresh()
}

func buildDrives() {
	buildJump()
	buildFuel()
	buildManeuver()
	buildPower()
}

func drivesTonsUsed() int {
	return int(float32(StarShip.jumpTons) + float32(StarShip.maneuverTons) + float32(StarShip.powerTons) + float32(StarShip.bridge) + float32(StarShip.jumpFuel))
}
