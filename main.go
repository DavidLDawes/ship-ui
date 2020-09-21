package main

import (
	"strconv"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func buildDetails() {
	buildShip()
	buildDrives()
	buildWeapons()
	buildBerths()
	buildVehicles()
	buildTotal()
}

var weaponSettings *widget.Form

func main() {

	buildDetails()
	weaponsInit()

	a := app.New()
	w := a.NewWindow("Designer")

	tlSelect = widget.NewSelect(techLevels, techLevelChanged)
	tlSelect.SetSelected("F")

	tonsSelect = widget.NewSelect(tons, tonsChanged)
	tonsSelect.SetSelected("200")

	jumpSelect = widget.NewSelect(engineLevel, jumpChanged)
	jumpSelect.SetSelected("2")

	maneuverSelect = widget.NewSelect(engineLevel, maneuverChanged)
	maneuverSelect.SetSelected("2")

	powerSelect = widget.NewSelect(engineLevel, powerChanged)
	powerSelect.SetSelected("2")

	shipSettings := widget.NewForm(
		widget.NewFormItem("Tech Level", tlSelect),
		widget.NewFormItem("tons", tonsSelect),
		widget.NewFormItem("Jump", jumpSelect),
		widget.NewFormItem("Maneuver", maneuverSelect),
		widget.NewFormItem("Power", powerSelect),
	)

	shipDetails := widget.NewVBox(
		&detailTons,
		&detailJump,
		&detailJumpFuel,
		&detailManeuver,
		&detailPower,
		&detailBridge,
		&detailComputer,
		&detailHardPoints,
		&detailTotal,
	)

	missileSelect = widget.NewSelect(weaponLevel, missileChanged)
	missileSelect.SetSelected("0")
	beamSelect = widget.NewSelect(weaponLevel, beamChanged)
	beamSelect.SetSelected("0")
	pulseSelect = widget.NewSelect(weaponLevel, pulseChanged)
	pulseSelect.SetSelected("0")
	fusionSelect = widget.NewSelect(weaponLevel, fusionChanged)
	fusionSelect.SetSelected("0")
	plasmaSelect = widget.NewSelect(weaponLevel, plasmaChanged)
	plasmaSelect.SetSelected("0")
	sandSelect = widget.NewSelect(weaponLevel, sandChanged)
	sandSelect.SetSelected("0")
	particleSelect = widget.NewSelect(weaponLevel, particleChanged)
	particleSelect.SetSelected("0")

	weaponSettings = widget.NewForm(
		widget.NewFormItem("Missile", missileSelect),
		widget.NewFormItem("Beam", beamSelect),
		widget.NewFormItem("Pulse", pulseSelect),
		widget.NewFormItem("Fusion", fusionSelect),
		widget.NewFormItem("Sand", sandSelect),
		widget.NewFormItem("Plasma", plasmaSelect),
		widget.NewFormItem("Accelerators", particleSelect),
	)

	weaponDetails := widget.NewVBox(
		&detailMissile,
		&detailBeam,
		&detailPulse,
		&detailFusion,
		&detailSand,
		&detailPlasma,
		&detailParticle,
	)

	lowLevel := make([]string, 101)
	for i := 0; i < 101; i++ {
		lowLevel[i] = strconv.Itoa(i)
	}

	lowBerthSelect = widget.NewSelect(lowLevel, lowBerthsChanged)
	lowBerthSelect.SetSelected("0")

	emergencyLowSelect = widget.NewSelect(lowLevel, emergencyLowChanged)
	emergencyLowSelect.SetSelected("1")

	adjustSlider()

	berthSettings := widget.NewForm(
		widget.NewFormItem("Staterooms", stateroomSlider),
		widget.NewFormItem("Low Berths", lowBerthSelect),
		widget.NewFormItem("Emergency Low Berths", emergencyLowSelect),
	)

	berthDetails := widget.NewVBox(
		&detailStaterooms,
		&detailLowBerths,
		&detailEmergencyLow,
		&detailCommandCrew,
		&detailBridgeCrew,
		&detailEngCrew,
		&detailGunCrew,
		&detailStewardCrew,
	)

	atvWheelSelect = widget.NewSelect(weaponLevel, atvWheelChanged)
	atvTrackSelect = widget.NewSelect(weaponLevel, atvTrackChanged)
	airRaftSelect = widget.NewSelect(weaponLevel, airRaftChanged)
	speederSelect = widget.NewSelect(weaponLevel, speederChanged)
	gCarrierSelect = widget.NewSelect(weaponLevel, gCarrierChanged)
	launchSelect = widget.NewSelect(weaponLevel, launchChanged)
	shipsBoatSelect = widget.NewSelect(weaponLevel, shipsBoatChanged)
	pinnaceSelect = widget.NewSelect(weaponLevel, pinnaceChanged)
	cutterSelect = widget.NewSelect(weaponLevel, cutterChanged)
	slowBoatSelect = widget.NewSelect(weaponLevel, slowBoatChanged)
	slowPinnaceSelect = widget.NewSelect(weaponLevel, slowPinnaceChanged)
	shuttleSelect = widget.NewSelect(weaponLevel, shuttleChanged)
	ltFigherSelect = widget.NewSelect(weaponLevel, ltFigherChanged)
	medFigherSelect = widget.NewSelect(weaponLevel, medFighterChanged)
	hvyFigherSelect = widget.NewSelect(weaponLevel, hvyFighterChanged)

	vehicleSettings := widget.NewForm(
		widget.NewFormItem("ATV, Wheeled", atvWheelSelect),
		widget.NewFormItem("ATV, Tracked", atvTrackSelect),
		widget.NewFormItem("Air/Raft", airRaftSelect),
		widget.NewFormItem("Speeder", speederSelect),
		widget.NewFormItem("GCarrier", gCarrierSelect),
		widget.NewFormItem("Launch", launchSelect),
		widget.NewFormItem("Ship's Boat", shipsBoatSelect),
		widget.NewFormItem("Pinnace", pinnaceSelect),
		widget.NewFormItem("Cutter", cutterSelect),
		widget.NewFormItem("Slow Boat", slowBoatSelect),
		widget.NewFormItem("Slow Pinnace", slowPinnaceSelect),
		widget.NewFormItem("Shuttle", shuttleSelect),
		widget.NewFormItem("Light Fighter", ltFigherSelect),
		widget.NewFormItem("Medium Fighter", medFigherSelect),
		widget.NewFormItem("Heavy Fighter", hvyFigherSelect),
	)

	vehicleDetails := widget.NewVBox(
		&detailSurfaceVehicles,
		&detailUtilityVehicles,
		&detailHighEndVehicles,
	)

	adjustSlider()

	ui := widget.NewVBox(widget.NewHBox(widget.NewLabel("Drives"), shipSettings, shipDetails, widget.NewLabel("Weapons"), weaponSettings, weaponDetails),
		widget.NewLabel("Berths and Crew"), widget.NewHBox(berthSettings, berthDetails, vehicleSettings, vehicleDetails),
	)
	w.SetContent(ui)

	w.ShowAndRun()
}
