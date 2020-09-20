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
}

var StarShip = shipDetails{
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
	}
	buildTons()
	buildBridge()
	BuildTotal()
}

func buildTons() {
	detailTons.SetText(fmt.Sprintf("Tons: %d", StarShip.tons))
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
func BuildTotal() {
	total := shipTonsUsed() + drivesTonsUsed()

	detailTotal.SetText(fmt.Sprintf("Tons remaining: %2.1f", total))
	detailTotal.Refresh()
}

func buildShip() {
	buildTons()
	buildBridge()
	buildHardPoints()
	BuildTotal()
}

func shipTonsUsed() int {
	return int(float32(StarShip.bridge) + float32(StarShip.computer) + float32(StarShip.hardpoints))
}
