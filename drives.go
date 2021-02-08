package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

var detailJump *widget.Label = widget.NewLabel("")
var detailJumpFuel *widget.Label = widget.NewLabel("")
var detailManeuver *widget.Label = widget.NewLabel("")
var detailPower *widget.Label = widget.NewLabel("")

var jumpSelect *widget.Select
var maneuverSelect *widget.Select
var powerSelect *widget.Select

func drivesInit() {
	jumpSelect = widget.NewSelect(engineLevel, nothing)
	maneuverSelect = widget.NewSelect(engineLevel, nothing)
	powerSelect = widget.NewSelect(engineLevel, nothing)
}

func drivesSelectsInit() {
	jumpSelect.SetSelected("2")
	maneuverSelect.SetSelected("2")
	powerSelect.SetSelected("2")

	jumpSelect.OnChanged = jumpChanged
	maneuverSelect.OnChanged = maneuverChanged
	powerSelect.OnChanged = powerChanged
}

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

var jbase = [6]float32{
	2.0, 1.75, 1.5, 1.0, 0.66666, 0.5,
}
var jinc = [6]float32{
	1.0, .8, .66666, 0.5, 0.33333, 0.25,
}

func jumpRate() float32 {
	return getDiscount() * (jbase[StarShip.tlOffset] + jinc[StarShip.tlOffset]*(float32(StarShip.jump)-1.0)) / 100.0
}

var jfbase = [6]float32{
	10.0, 8.0, 6.0, 5.0, 4.0, 2.5,
}
var jfinc = [6]float32{
	10.0, 7.0, 6.0, 4.0, 2.5, 1.75,
}

func jumpFuelRate() float32 {
	return getFuelDiscount() * (jfbase[StarShip.tlOffset] + jfinc[StarShip.tlOffset]*(float32(StarShip.jump)-1.0)) / 100.0
}

var mbase = [6]float32{
	3.0, 1.5, 1.25, 1.0, 0.66666, 0.5,
}
var minc = [6]float32{
	4.0, 2.0, 1.66666, 1.5, 1.0, 0.75,
}

func maneuverRate() float32 {
	return getDiscount() * (mbase[StarShip.tlOffset] + minc[StarShip.tlOffset]*(float32(StarShip.maneuver)-1.0)) / 100.0
}

var pbase = [6]float32{
	1.0, 0.7, 0.5, 0.4, 0.33333, 0.25,
}
var pinc = [6]float32{
	1.0, .7, 0.55, 0.4, 0.25, 0.175,
}

func powerRate() float32 {
	return getDiscount() * (pbase[StarShip.tlOffset] + pinc[StarShip.tlOffset]*(float32(StarShip.power)-1.0)) / 100.0
}

func buildJump() {
	StarShip.jumpTons = float32(StarShip.tons) * jumpRate() * armor()
	detailJump.SetText(fmt.Sprintf("Jump: %d, tons: %2.1f", StarShip.jump, StarShip.jumpTons))
	detailJump.Refresh()
	detailComputer.SetText(fmt.Sprintf("computer %d: %d tons", StarShip.jump, int(armor()*float32(computer[StarShip.jump-1])+.9999)))
	detailComputer.Refresh()
	setEngineers()
	refreshEngineeringCrew()
}
func buildManeuver() {
	StarShip.maneuverTons = float32(StarShip.tons) * maneuverRate() * armor()
	detailManeuver.SetText(fmt.Sprintf("Maneuver: %d, tons: %2.1f", StarShip.maneuver, StarShip.maneuverTons))
	detailManeuver.Refresh()
	setEngineers()
	refreshEngineeringCrew()
}
func buildPower() {
	StarShip.powerTons = float32(StarShip.tons) * powerRate() * armor()
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

func nothing(value string) {
}
func nothingBool(value bool) {
}

func getDiscount() (discount float32) {
	discount = float32(1.0)
	if StarShip.tlOffset == 3 {
		if StarShip.tons > 99999 {
			discount = 0.66666
		} else if StarShip.tons > 9999 {
			discount = 0.75
		} else if StarShip.tons > 999 {
			discount = 0.9
		}
	} else if StarShip.tlOffset == 4 {
		discount = 0.9333
		if StarShip.tons > 99999 {
			discount = 0.5
		} else if StarShip.tons > 9999 {
			discount = 0.66666
		} else if StarShip.tons > 999 {
			discount = 0.8666
		}
	} else if StarShip.tlOffset == 5 {
		discount = 0.86666
		if StarShip.tons > 99999 {
			discount = 0.33333
		} else if StarShip.tons > 9999 {
			discount = 0.5
		} else if StarShip.tons > 999 {
			discount = 0.75
		}
	}
	return discount
}

func getFuelDiscount() (discount float32) {
	discount = float32(1.0)
	if StarShip.tlOffset == 3 {
		discount = 0.6666
	} else if StarShip.tlOffset == 4 {
		discount = .5
	} else if StarShip.tlOffset == 5 {
		discount = .363636
	}
	return discount
}