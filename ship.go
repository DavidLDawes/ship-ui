package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type shipDetails struct {
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
}

var StarShip = &shipDetails{
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
}

var detailTons widget.Label
var detailBridge widget.Label
var detailComputer widget.Label
var detailHardPoints widget.Label
var detailTotal widget.Label

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
}

func buildTons() {
	detailTons.SetText(fmt.Sprintf("tons: %d", StarShip.tons))
	detailTons.Refresh()
}
func buildBridge() {
	bridgeTons := int(float32(StarShip.tons) * .02)
	if bridgeTons < 20 {
		bridgeTons = 20
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
	total := StarShip.tons - shipTonsUsed() - drivesTonsUsed() - weaponsTonsUsed() - berthsTonsUsed()
	StarShip.remaining = float32(total)
	detailTotal.SetText(fmt.Sprintf("tons remaining: %d", total))
	detailTotal.Refresh()
	save := berths.staterooms
	berths.staterooms = getTotalCrew()
	adjustSlider()

	if save > getTotalCrew() {
		if float64(save) < stateroomSlider.Max {
			berths.staterooms = save
			stateroomSlider.Value = float64(save)
		} else {
			berths.staterooms = int(stateroomSlider.Max)
			stateroomSlider.Value = stateroomSlider.Max
		}
	}
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
