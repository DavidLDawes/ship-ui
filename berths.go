package main

import (
	"fmt"
	"math"
	"strconv"

	"fyne.io/fyne/widget"
)

type berthDetails struct {
	staterooms   int
	lowBerths    int
	emergencylow int
	pilots       int
	engineer     int
	stewards     int
	navigator    int
	medic        int
	gunners      int
	exec         int
	command      int
	computer     int
	comms        int
	support      int
	security     int
	service      int
}

var berths = berthDetails{
	staterooms:   4,
	lowBerths:    0,
	emergencylow: 1,
	pilots:       1,
	engineer:     1,
	stewards:     1,
	navigator:    1,
	medic:        0,
	gunners:      0,
	exec:         0,
	command:      0,
	computer:     0,
	comms:        0,
	support:      0,
	security:     0,
	service:      0,
}

var detailStaterooms widget.Label
var detailLowBerths widget.Label
var detailEmergencyLow widget.Label
var detailCommandCrew widget.Label
var detailBridgeCrew widget.Label
var detailEngCrew widget.Label
var detailGunCrew widget.Label
var detailStewardCrew widget.Label

var stateroomSlider *widget.Slider
var lowBerthSelect *widget.Select
var emergencyLowSelect *widget.Select

var ignoreBerthChanges = false

func stateroomChanged(rooms float64) {
	rooms = math.Floor(rooms + .999)
	if int(rooms) < getTotalCrew() {
		rooms = float64(getTotalCrew())
		ignoreBerthChanges = true
		stateroomSlider.Value = rooms
		ignoreBerthChanges = false
	}
	berths.staterooms = int(rooms)
	buildStaterooms()
	buildCrew()
	buildTotal()
}

func lowBerthsChanged(value string) {
	if !ignoreBerthChanges {
		low, err := strconv.Atoi(value)
		if err == nil {
			if low > -1 {
				berths.lowBerths = low
				buildLowBerths()
				buildCrew()
				buildTotal()
			}
		}
	}
}

func emergencyLowChanged(value string) {
	if !ignoreBerthChanges {
		elow, err := strconv.Atoi(value)
		if err == nil {
			if elow > -1 {
				berths.emergencylow = elow
				buildEmergencyLow()
				buildCrew()
				buildTotal()
			}
		}
	}
}

func buildStaterooms() {
	ignoreBerthChanges = true
	detailStaterooms.SetText(fmt.Sprintf("Staterooms: %d, tons: %d", berths.staterooms, 4*berths.staterooms))
	ignoreBerthChanges = false
	detailStaterooms.Refresh()
}
func buildLowBerths() {
	ignoreBerthChanges = true
	detailLowBerths.SetText(fmt.Sprintf("Low berths: %d, tons: %d", berths.lowBerths, (1+berths.lowBerths)/2))
	ignoreBerthChanges = false
	detailLowBerths.Refresh()
}
func buildEmergencyLow() {
	ignoreBerthChanges = true
	detailEmergencyLow.SetText(fmt.Sprintf("Emergency low berths: %d, tons: %d", berths.emergencylow, berths.emergencylow))
	ignoreBerthChanges = false
	detailLowBerths.Refresh()
}
func buildCrew() {
	berths.pilots = 1
	setEngineers()
	berths.gunners = countWeapons()
	berths.service = int(StarShip.tons/1000) * 2
	if StarShip.tons > 1000 {
		berths.command = 1
		berths.exec = 1
		berths.computer = 1
		berths.comms = 1
		berths.navigator = 2
		berths.medic = 1
		berths.support = 4
		berths.security = StarShip.tons / 333
		if StarShip.tons > 20000 {
			berths.support = StarShip.tons / 2000
			if berths.support < 4 {
				berths.support = 4
			}
		}
	} else {
		berths.command = 0
		berths.exec = 0
		berths.computer = 0
		berths.comms = 0
		berths.navigator = 1
		berths.medic = 0
		berths.support = 0
		berths.security = 0
	}
	berths.stewards = 0
	berths.stewards = (6 + getTotalCrew()) / 7

	cmdCrew := ""
	if berths.command > 0 {
		cmdCrew = "1 Commander, "
	}

	if berths.exec > 0 {
		cmdCrew = cmdCrew + fmt.Sprintf("%d Exec, ", berths.exec)
	}

	if berths.computer > 0 {
		cmdCrew = cmdCrew + fmt.Sprintf("%d Computer, ", berths.computer)
	}

	if berths.comms > 0 {
		cmdCrew = cmdCrew + fmt.Sprintf("%d Comms, ", berths.comms)
	}
	detailCommandCrew.SetText(cmdCrew)
	detailCommandCrew.Refresh()

	brdgCrew := fmt.Sprintf("%d Pilot, ", berths.pilots)
	if berths.navigator > 0 {
		brdgCrew = brdgCrew + fmt.Sprintf("%d Nav, ", berths.navigator)
	}
	if berths.medic > 0 {
		brdgCrew = brdgCrew + fmt.Sprintf("%d Medic, ", berths.medic)
	}
	detailBridgeCrew.SetText(brdgCrew)
	detailBridgeCrew.Refresh()

	refreshEngineeringCrew()

	if berths.security > 0 {
		if berths.gunners > 0 {
			detailGunCrew.SetText(fmt.Sprintf("%d Gunners, %d Security", berths.gunners, berths.security))
		} else {
			detailGunCrew.SetText(fmt.Sprintf("%d Security", berths.security))
		}
	} else {
		if berths.gunners > 0 {
			detailGunCrew.SetText(fmt.Sprintf("%d Gunners", berths.gunners))
		} else {
			detailGunCrew.SetText(fmt.Sprintf("No Gunners, No Security"))
		}
	}
	detailGunCrew.Refresh()

	if getTotalCrew() > 120 {
		berths.medic = (119 + berths.staterooms) / 120
	}
	berths.stewards = 0
	berths.stewards = (6 + getTotalCrew()) / 7
	if getTotalCrew() > 120 {
		berths.medic = (119 + berths.staterooms) / 120
	}
	berths.stewards = (6 + getTotalCrew()) / 7
	if getTotalCrew() < berths.staterooms {
		berths.stewards = (6 + berths.staterooms) / 7
	}

	if berths.staterooms < getTotalCrew() {
		berths.staterooms = getTotalCrew()
		buildStaterooms()
		stateroomSlider.Value = float64(berths.staterooms)
	}

	if berths.support > 0 {
		detailStewardCrew.SetText(fmt.Sprintf("%d Stewards, %d Support", berths.stewards, berths.support))
	} else {
		detailStewardCrew.SetText(fmt.Sprintf("%d Stewards", berths.stewards))
	}
	detailStewardCrew.Refresh()
}
func buildBerths() {
	buildStaterooms()
	buildLowBerths()
	buildCrew()
}

func setEngineers() {
	tmp := int((StarShip.jumpTons + StarShip.maneuverTons + StarShip.powerTons + 99) / 100)
	berths.engineer = tmp
}

func refreshEngineeringCrew() {
	if berths.service > 0 {
		detailEngCrew.SetText(fmt.Sprintf("%d Engineers, %d Service", berths.engineer, berths.service))
	} else {
		detailEngCrew.SetText(fmt.Sprintf("%d Engineers", berths.engineer))
	}
	detailEngCrew.Refresh()
}

func refreshPilots() {
	berths.pilots = 1 + countVehicles()
	detailBridgeCrew.Refresh()
}

func adjustSlider() {
	maxStaterooms := float64(StarShip.remaining) / 4.0
	stateroomSlider = widget.NewSlider(float64(getTotalCrew()), maxStaterooms)
	stateroomSlider.OnChanged = stateroomChanged
	tmp := getTotalCrew()
	if tmp > -1 {
		stateroomSlider.Value = float64(getTotalCrew())
	}
}

func getTotalCrew() int {
	refreshPilots()
	return berths.engineer + berths.pilots + berths.gunners + berths.medic + berths.stewards + berths.navigator + berths.exec + berths.command + berths.computer + berths.comms + berths.security + berths.support + berths.service
}

func berthsTonsUsed() int {
	return 4*berths.staterooms + (berths.lowBerths+1)/2
}
