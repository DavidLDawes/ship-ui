package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type shipCountsAndTons struct {
	tons         int
	tl           string
	tlOffset     int
	jump         int
	jumpTons     float32
	jumpFuel     int
	maneuver     int
	maneuverTons float32
	power        int
	powerTons    float32
	bridge       int
	computer     int
	hardpoints   int
	remaining    float32
	armored      bool
	robo         bool
}

// StarShip originally the definition, now a subset
var StarShip = &shipCountsAndTons{
	tons:         200,
	tl:           "F",
	tlOffset:     0,
	jump:         2,
	jumpTons:     12,
	maneuver:     2,
	maneuverTons: 8,
	power:        2,
	powerTons:    8,
	bridge:       40,
	computer:     2,
	hardpoints:   2,
	remaining:    83,
	armored:      false,
	robo:         false,
}

var tonsSelect *widget.Select
var armoredSelect *widget.Check
var roboSelect *widget.Check

var detailTons *widget.Label = widget.NewLabel("")
var detailBridge *widget.Label = widget.NewLabel("")
var detailComputer *widget.Label = widget.NewLabel("")
var detailHardPoints *widget.Label = widget.NewLabel("")
var detailTotal *widget.Label = widget.NewLabel("")

var shipSettings *widget.Form

var shipDetails *widget.Box

func shipInit() {
	tonsSelect = widget.NewSelect(tons, tonsChanged)
	tonsSelect.Selected = "200"

	jumpSelect = widget.NewSelect(engineLevel, jumpChanged)
	jumpSelect.SetSelected("2")

	maneuverSelect = widget.NewSelect(engineLevel, maneuverChanged)
	maneuverSelect.SetSelected("2")

	powerSelect = widget.NewSelect(engineLevel, powerChanged)
	powerSelect.SetSelected("2")

	armoredSelect = widget.NewCheck("Armored bulkheads", armoredChanged)
	armoredSelect.Checked = false

	roboSelect = widget.NewCheck("Robotic crew", roboChanged)
	roboSelect.Checked = false

	shipSettings = widget.NewForm(
		widget.NewFormItem("Tech Level", tlSelect),
		widget.NewFormItem("tons", tonsSelect),
		widget.NewFormItem("Jump", jumpSelect),
		widget.NewFormItem("Maneuver", maneuverSelect),
		widget.NewFormItem("Power", powerSelect),
		widget.NewFormItem("Armor", armoredSelect),
		widget.NewFormItem("Robots", roboSelect),
	)

	shipDetails = widget.NewVBox(
		detailTons,
		detailJump,
		detailJumpFuel,
		detailManeuver,
		detailPower,
		detailBridge,
		detailComputer,
		detailHardPoints,
		detailTotal,
	)
}

func armoredChanged(armored bool) {
	StarShip.armored = armored
	tonsChanged(strconv.Itoa(StarShip.tons))
}
func roboChanged(robots bool) {
	StarShip.robo = robots
	buildBerths()
}
func tonsChanged(value string) {
	tons, err := strconv.Atoi(value)
	if err == nil {
		StarShip.tons = tons
		StarShip.hardpoints = StarShip.tons / 100

		if countWeapons() > StarShip.hardpoints {
			weapons.missile = 0
			ignoreMissile = true
			missileSelect.SetSelected("0")
			ignoreMissile = false
			buildMissile()
			if countWeapons() > StarShip.hardpoints {
				weapons.beam = 0
				ignoreBeam = true
				beamSelect.SetSelected("0")
				ignoreBeam = false
				buildBeam()
				if countWeapons() > StarShip.hardpoints {
					weapons.pulse = 0
					ignorePulse = true
					pulseSelect.SetSelected("0")
					ignorePulse = false
					buildPulse()
					if countWeapons() > StarShip.hardpoints {
						weapons.plasma = 0
						ignorePlasma = true
						plasmaSelect.SetSelected("0")
						ignorePlasma = false
						buildPlasma()
						if countWeapons() > StarShip.hardpoints {
							weapons.fusion = 0
							ignoreFusion = true
							fusionSelect.SetSelected("0")
							ignoreFusion = false
							buildFusion()
							if countWeapons() > StarShip.hardpoints {
								weapons.accelerator = 0
								ignoreParticle = true
								particleSelect.SetSelected("0")
								ignoreParticle = false
								buildParticle()
								if countWeapons() > StarShip.hardpoints {
									weapons.sandcaster = 0
									ignoreSand = true
									sandSelect.SetSelected("0")
									ignoreSand = false
									buildSand()
								}
							}
						}
					}
				}
			}
		}
	}
	buildDrives()
	buildBridge()
	buildHardPoints()
	buildWeapons()
	setEngineers()
	refreshEngineeringCrew()
	buildCrew()
	buildTons()
	buildTotal()
	adjustSlider()
}

func buildTons() {
	detailTons.SetText(fmt.Sprintf("tons: %d", StarShip.tons))
	detailTons.Refresh()
}
func buildBridge() {
	bridgeTons := int(float32(StarShip.tons) * .02 * armor())
	if bridgeTons < 20 {
		bridgeTons = int(armor()*float32(20) + .9999)
	}
	StarShip.bridge = bridgeTons
	detailBridge.SetText(fmt.Sprintf("Bridge: %d", bridgeTons))
	detailBridge.Refresh()
}
func buildHardPoints() {
	StarShip.hardpoints = StarShip.tons / 100
	detailHardPoints.SetText(fmt.Sprintf("Hardpoints: %d", StarShip.hardpoints))
	detailHardPoints.Refresh()
}
func buildTotal() {
	StarShip.remaining = remainingTons()
	detailTotal.SetText(fmt.Sprintf("Tons remaining: %2.1f", StarShip.remaining))
	if berths.staterooms < getTotalCrew() {
		berths.staterooms = getTotalCrew()
		StarShip.remaining = remainingTons()
		detailTotal.SetText(fmt.Sprintf("Tons remaining: %2.1f", StarShip.remaining))
	} else {
		if StarShip.remaining < 0 {
			berths.staterooms = getTotalCrew()
			StarShip.remaining = remainingTons()
			detailTotal.SetText(fmt.Sprintf("Tons remaining: %2.1f", StarShip.remaining))
		}
	}
	detailTotal.Refresh()
}

func buildShip() {
	buildTons()
	buildBridge()
	buildHardPoints()
	buildTotal()
}

func shipTonsUsed() int {
	return int(float32(StarShip.bridge) + float32(StarShip.computer) + float32(StarShip.hardpoints))
}

func remainingTons() float32 {
	remaining := StarShip.tons - shipTonsUsed() - drivesTonsUsed() - weaponsTonsUsed() - berthsTonsUsed() - vehicleTonsUsed()
	return float32(remaining)

}
