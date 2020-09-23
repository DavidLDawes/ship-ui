package main

import (
	"fmt"
	"math"
	"strconv"

	"fyne.io/fyne/widget"
)

type berthCounts struct {
	staterooms   int
	lowBerths    int
	emergencylow int
	pilots       int
	engineer     int
	stewards     int
	roboStewards int
	navigator    int
	medic        int
	gunners      int
	roboGunners  int
	exec         int
	command      int
	computer     int
	comms        int
	support      int
	roboSupport  int
	security     int
	roboSecurity int
	service      int
	roboService  int
}

var berths = berthCounts{
	staterooms:   4,
	lowBerths:    0,
	emergencylow: 1,
	pilots:       1,
	engineer:     1,
	stewards:     1,
	roboStewards: 0,
	navigator:    1,
	medic:        0,
	gunners:      0,
	roboGunners:  0,
	exec:         0,
	command:      0,
	computer:     0,
	comms:        0,
	support:      0,
	roboSupport:  0,
	security:     0,
	roboSecurity: 0,
	service:      0,
	roboService:  0,
}

var detailStaterooms *widget.Label = widget.NewLabel("")
var detailLowBerths *widget.Label = widget.NewLabel("")
var detailEmergencyLow *widget.Label = widget.NewLabel("")
var detailCommandCrew *widget.Label = widget.NewLabel("")
var detailBridgeCrew *widget.Label = widget.NewLabel("")
var detailEngCrew *widget.Label = widget.NewLabel("")
var detailGunCrew *widget.Label = widget.NewLabel("")
var detailStewardCrew *widget.Label = widget.NewLabel("")

var stateroomSlider *widget.Slider
var lowBerthSelect *widget.Select
var emergencyLowSelect *widget.Select

var berthSettings *widget.Form
var berthDetails *widget.Box

var ignoreBerthChanges = false

func berthsInit() {
	stateroomSlider = widget.NewSlider(4.0, 28.0)
	stateroomSlider.Value = 4.0
	stateroomSlider.OnChanged = stateroomChanged

	lowLevel := make([]string, 401)
	for i := 0; i < 401; i++ {
		lowLevel[i] = strconv.Itoa(i)
	}

	lowBerthSelect = widget.NewSelect(lowLevel, lowBerthsChanged)
	emergencyLowSelect = widget.NewSelect(lowLevel, emergencyLowChanged)

	berthSettings = widget.NewForm(
		widget.NewFormItem("Staterooms", stateroomSlider),
		widget.NewFormItem("Low Berths", lowBerthSelect),
		widget.NewFormItem("Emergency Low Berths", emergencyLowSelect),
	)

	adjustSlider()

	berthDetails = widget.NewVBox(
		detailStaterooms,
		detailLowBerths,
		detailEmergencyLow,
		detailCommandCrew,
		detailBridgeCrew,
		detailEngCrew,
		detailGunCrew,
		detailStewardCrew,
	)
}

func berthsSelectsInit() {
	lowBerthSelect.SetSelected("0")
	emergencyLowSelect.SetSelected("0")
}

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
	detailStaterooms.SetText(fmt.Sprintf("Staterooms: %d, tons: %d", berths.staterooms, int(armor()*float32(4*berths.staterooms)+.9999)))
	ignoreBerthChanges = false
	detailStaterooms.Refresh()
}
func buildLowBerths() {
	ignoreBerthChanges = true
	detailLowBerths.SetText(fmt.Sprintf("Low berths: %d, tons: %d", berths.lowBerths, int(armor()*float32(berths.lowBerths/2)+.9999)))
	ignoreBerthChanges = false
	detailLowBerths.Refresh()
}
func buildEmergencyLow() {
	ignoreBerthChanges = true
	detailEmergencyLow.SetText(fmt.Sprintf("Emergency low berths: %d, tons: %d", berths.emergencylow, int(armor()*float32(berths.emergencylow)+.9999)))
	ignoreBerthChanges = false
	detailLowBerths.Refresh()
}
func buildCrew() {
	berths.pilots = 1
	setEngineers()
	if StarShip.robo {
		berths.roboGunners = int(float32(countWeapons()) * .75)
		berths.gunners = countWeapons() - berths.roboGunners
	} else {
		berths.roboGunners = 0
		berths.gunners = countWeapons()
	}
	if StarShip.robo {
		berths.roboService = int(float32(2*StarShip.tons) * 0.75 / 1000)
		berths.service = int(StarShip.tons/1000)*2 - berths.roboService
	} else {
		berths.roboService = 0
		berths.service = int(StarShip.tons/1000) * 2
	}
	if StarShip.tons > 1000 {
		berths.command = 1
		berths.exec = 1
		berths.computer = 1
		berths.comms = 1
		berths.navigator = 2
		berths.medic = 1
		if StarShip.robo {
			berths.roboSupport = 3
			berths.support = 1
		} else {
			berths.roboSupport = 0
			berths.support = 4
		}
		if StarShip.robo {
			berths.roboSecurity = int(float32(StarShip.tons) * .75 / 333.3333)
			berths.security = StarShip.tons/333 - berths.roboSecurity
		} else {
			berths.roboSecurity = 0
			berths.security = StarShip.tons / 333
		}
		if StarShip.tons > 20000 {
			support := StarShip.tons / 2000
			if support < 4 {
				support = 4
			}
			if StarShip.robo {
				berths.roboSupport = int(float32(support) * 0.75)
				berths.support = support - berths.roboSupport
			} else {
				berths.roboSupport = 0
				berths.support = support
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
	setStewards()
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
			if StarShip.robo {
				detailGunCrew.SetText(fmt.Sprintf("%d Gunners, %d RoboGunners %d Security, %d RoboSecurity", berths.gunners, berths.roboGunners, berths.security, berths.roboSecurity))
			} else {
				detailGunCrew.SetText(fmt.Sprintf("%d Gunners, %d Security", berths.gunners, berths.security))
			}
		} else {
			if StarShip.robo {
				detailGunCrew.SetText(fmt.Sprintf("%d Security, %d RoboSecurity", berths.security, berths.roboSecurity))
			} else {
				detailGunCrew.SetText(fmt.Sprintf("%d Security", berths.security))
			}
		}
	} else {
		if berths.gunners > 0 {
			if StarShip.robo {
				detailGunCrew.SetText(fmt.Sprintf("%d Gunners, %d RoboGunners", berths.gunners, berths.roboGunners))
			} else {
				detailGunCrew.SetText(fmt.Sprintf("%d Gunners", berths.gunners))
			}
		} else {
			detailGunCrew.SetText(fmt.Sprintf("No Gunners, No Security"))
		}
	}
	detailGunCrew.Refresh()

	if getTotalCrew() > 120 {
		berths.medic = (119 + berths.staterooms) / 120
	}
	setStewards()
	if getTotalCrew() > 120 {
		berths.medic = (119 + berths.staterooms) / 120
	}
	setStewards()

	if berths.staterooms < getTotalCrew() {
		berths.staterooms = getTotalCrew()
		buildStaterooms()
		stateroomSlider.Value = float64(berths.staterooms)
	}

	if berths.support > 0 {
		if StarShip.robo {
			detailStewardCrew.SetText(fmt.Sprintf("%d Stewards, %d RobotStewards, %d Support, %d RoboSupport", berths.stewards, berths.roboStewards, berths.support, berths.roboSupport))
		} else {
			detailStewardCrew.SetText(fmt.Sprintf("%d Stewards, %d Support", berths.stewards, berths.support))
		}
	} else {
		if StarShip.robo {
			detailStewardCrew.SetText(fmt.Sprintf("%d Stewards, %d RoboStewards", berths.stewards, berths.roboStewards))
		} else {
			detailStewardCrew.SetText(fmt.Sprintf("%d Stewards", berths.stewards))
		}
	}
	detailStewardCrew.Refresh()
}
func buildBerths() {
	buildStaterooms()
	buildLowBerths()
	buildCrew()
}

func setEngineers() {
	tmp := int((StarShip.jumpTons + StarShip.maneuverTons + StarShip.powerTons) / (armor() * 100.0))
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
	maxStaterooms := float64(remainingTons() / 4.0)
	minStaterooms := getTotalCrew()
	stateroomSlider.Min = float64(minStaterooms)
	stateroomSlider.Max = float64(maxStaterooms)
}

func getTotalCrew() int {
	refreshPilots()
	return berths.engineer + berths.pilots + berths.gunners + berths.medic + berths.stewards + berths.navigator + berths.exec + berths.command + berths.computer + berths.comms + berths.security + berths.support + berths.service
}

func getTotalRobots() int {
	return berths.roboGunners + berths.roboSecurity + berths.roboService + berths.roboStewards + berths.roboSupport
}

func berthsTonsUsed() int {
	return 4*berths.staterooms + (berths.lowBerths+1)/2 + (getTotalRobots()+15)/16
}

func setStewards() {
	berths.stewards = 0
	berths.roboStewards = 0
	if StarShip.robo {
		tot := (6 + getTotalCrew()) / 7
		berths.roboStewards = int(float32(tot) * 0.75)
		berths.stewards = tot - berths.roboStewards
	} else {
		berths.stewards = (6 + getTotalCrew()) / 7
	}
}
