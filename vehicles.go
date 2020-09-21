package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type vehicleDetails struct {
	atvWheel    int
	atvTrack    int
	airRaft     int
	speeder     int
	gCarrier    int
	launch      int
	shipsBoat   int
	pinnace     int
	cutter      int
	slowBoat    int
	slowPinnace int
	shuttle     int
	ltFighter   int
	medFighter  int
	hvyFighter  int
}

var vehicles = vehicleDetails{
	atvWheel:    0,
	atvTrack:    0,
	airRaft:     0,
	speeder:     0,
	gCarrier:    0,
	launch:      0,
	shipsBoat:   0,
	pinnace:     0,
	cutter:      0,
	slowBoat:    0,
	slowPinnace: 0,
	shuttle:     0,
	ltFighter:   0,
	medFighter:  0,
	hvyFighter:  0,
}

var detailSurfaceVehicles widget.Label
var detailUtilityVehicles widget.Label
var detailHighEndVehicles widget.Label

var atvWheelSelect *widget.Select
var atvTrackSelect *widget.Select
var airRaftSelect *widget.Select
var speederSelect *widget.Select
var gCarrierSelect *widget.Select
var launchSelect *widget.Select
var shipsBoatSelect *widget.Select
var pinnaceSelect *widget.Select
var cutterSelect *widget.Select
var slowBoatSelect *widget.Select
var slowPinnaceSelect *widget.Select
var shuttleSelect *widget.Select
var ltFigherSelect *widget.Select
var medFigherSelect *widget.Select
var hvyFigherSelect *widget.Select

var ignorevehicles = false

func atvWheelChanged(value string) {
	if !ignorevehicles {
		atvw, err := strconv.Atoi(value)
		if err == nil {
			vehicles.atvWheel = atvw
			ignorevehicles = true
			buildSurface()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func atvTrackChanged(value string) {
	if !ignorevehicles {
		atvt, err := strconv.Atoi(value)
		if err == nil {
			vehicles.atvTrack = atvt
			ignorevehicles = true
			buildSurface()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func airRaftChanged(value string) {
	if !ignorevehicles {
		air, err := strconv.Atoi(value)
		if err == nil {
			vehicles.airRaft = air
			ignorevehicles = true
			buildSurface()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func speederChanged(value string) {
	if !ignorevehicles {
		spdr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.speeder = spdr
			ignorevehicles = true
			buildSurface()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func gCarrierChanged(value string) {
	if !ignorevehicles {
		gc, err := strconv.Atoi(value)
		if err == nil {
			vehicles.gCarrier = gc
			ignorevehicles = true
			buildSurface()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func launchChanged(value string) {
	if !ignorevehicles {
		launch, err := strconv.Atoi(value)
		if err == nil {
			vehicles.launch = launch
			ignorevehicles = true
			buildUtility()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func shipsBoatChanged(value string) {
	if !ignorevehicles {
		sboat, err := strconv.Atoi(value)
		if err == nil {
			vehicles.shipsBoat = sboat
			ignorevehicles = true
			buildUtility()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func pinnaceChanged(value string) {
	if !ignorevehicles {
		pinnace, err := strconv.Atoi(value)
		if err == nil {
			vehicles.pinnace = pinnace
			ignorevehicles = true
			buildUtility()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func cutterChanged(value string) {
	if !ignorevehicles {
		cutter, err := strconv.Atoi(value)
		if err == nil {
			vehicles.cutter = cutter
			ignorevehicles = true
			buildUtility()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func slowBoatChanged(value string) {
	if !ignorevehicles {
		sloboat, err := strconv.Atoi(value)
		if err == nil {
			vehicles.slowBoat = sloboat
			ignorevehicles = true
			buildUtility()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func slowPinnaceChanged(value string) {
	if !ignorevehicles {
		slopin, err := strconv.Atoi(value)
		if err == nil {
			vehicles.slowPinnace = slopin
			ignorevehicles = true
			buildUtility()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func shuttleChanged(value string) {
	if !ignorevehicles {
		shuttle, err := strconv.Atoi(value)
		if err == nil {
			vehicles.shuttle = shuttle
			ignorevehicles = true
			buildHighEnd()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func ltFigherChanged(value string) {
	if !ignorevehicles {
		lftr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.ltFighter = lftr
			ignorevehicles = true
			buildHighEnd()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func medFighterChanged(value string) {
	if !ignorevehicles {
		mftr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.medFighter = mftr
			ignorevehicles = true
			buildHighEnd()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func hvyFighterChanged(value string) {
	if !ignorevehicles {
		hftr, err := strconv.Atoi(value)
		if err == nil {
			vehicles.hvyFighter = hftr
			ignorevehicles = true
			buildHighEnd()
			buildCrew()
			buildTotal()
			ignorevehicles = false
		}
	}
}

func buildSurface() {
	surface := ""
	if vehicles.atvWheel > 0 {
		surface += fmt.Sprintf("%d ATV Wheeled, ", vehicles.atvWheel)
	}
	if vehicles.atvTrack > 0 {
		surface += fmt.Sprintf("%d ATV tracked, ", vehicles.atvTrack)
	}
	if vehicles.airRaft > 0 {
		surface += fmt.Sprintf("%d Air/Raft, ", vehicles.airRaft)
	}
	if vehicles.speeder > 0 {
		surface += fmt.Sprintf("%d Speeder, ", vehicles.speeder)
	}
	if vehicles.gCarrier > 0 {
		surface += fmt.Sprintf("%d GCarrier, ", vehicles.gCarrier)
	}
	detailSurfaceVehicles.SetText(surface)
	detailSurfaceVehicles.Refresh()
}

func buildUtility() {
	utility := ""
	if vehicles.launch > 0 {
		utility += fmt.Sprintf("%d Launch, ", vehicles.launch)
	}
	if vehicles.shipsBoat > 0 {
		utility += fmt.Sprintf("%d Ship's Boat, ", vehicles.shipsBoat)
	}
	if vehicles.pinnace > 0 {
		utility += fmt.Sprintf("%d Pinnace, ", vehicles.pinnace)
	}
	if vehicles.cutter > 0 {
		utility += fmt.Sprintf("%d Cutter, ", vehicles.cutter)
	}
	if vehicles.slowBoat > 0 {
		utility += fmt.Sprintf("%d Slow Boat, ", vehicles.slowBoat)
	}
	if vehicles.slowPinnace > 0 {
		utility += fmt.Sprintf("%d SLow Pinnace, ", vehicles.slowPinnace)
	}
	detailUtilityVehicles.SetText(utility)
	detailUtilityVehicles.Refresh()
}

func buildHighEnd() {
	highEnd := ""
	if vehicles.shuttle > 0 {
		highEnd += fmt.Sprintf("%d Shuttle, ", vehicles.shuttle)
	}
	if vehicles.ltFighter > 0 {
		highEnd += fmt.Sprintf("%d Light Fighter, ", vehicles.ltFighter)
	}
	if vehicles.medFighter > 0 {
		highEnd += fmt.Sprintf("%d Medium Fighter, ", vehicles.medFighter)
	}
	if vehicles.hvyFighter > 0 {
		highEnd += fmt.Sprintf("%d Heavy Fighter, ", vehicles.hvyFighter)
	}
	detailHighEndVehicles.SetText(highEnd)
	detailHighEndVehicles.Refresh()

}

func buildVehicles() {
	buildSurface()
	buildUtility()
	buildHighEnd()
	buildTotal()
}

func countVehicles() int {
	result := vehicles.atvWheel + vehicles.atvTrack + vehicles.airRaft + vehicles.speeder + vehicles.gCarrier + vehicles.launch + vehicles.shipsBoat + vehicles.pinnace + vehicles.cutter + vehicles.slowBoat + vehicles.slowPinnace + +vehicles.shuttle + vehicles.ltFighter + vehicles.medFighter + vehicles.hvyFighter
	return result
}

func vehicleTonsUsed() int {
	result := vehicles.atvWheel*10 + vehicles.atvTrack*10 + vehicles.airRaft*4 + vehicles.speeder*6 + vehicles.gCarrier*8 + vehicles.launch*20 + vehicles.shipsBoat*30 + vehicles.pinnace*40 + vehicles.cutter*80 + vehicles.slowBoat*30 + vehicles.slowPinnace*40 + vehicles.shuttle*95 + vehicles.ltFighter*10 + vehicles.medFighter*30 + vehicles.hvyFighter*50
	return result
}
