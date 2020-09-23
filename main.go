package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

var saveButton = widget.NewButton("Save", saveMe)
var loadButton = widget.NewButton("Load", loadMe)

func buildDetails() {
	buildShip()
	buildDrives()
	buildWeapons()
	buildBerths()
	buildVehicles()
	buildTotal()
}

func main() {

	a := app.New()
	w := a.NewWindow("Designer")

	tlSelect = widget.NewSelect(techLevels, nothing)
	tlSelect.SetSelected("F")
	tlSelect.OnChanged = techLevelChanged
	buildDetails()
	weaponsInit()
	berthsInit()
	drivesInit()
	shipInit()

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

	weaponsSelectInit()
	setVehicleDetails()
	drivesSelectsInit()
	berthsSelectsInit()
	adjustSlider()

	//	ui := widget.NewVBox(widget.NewHBox(widget.NewLabel("Drives"), shipSettings, shipDetails, widget.NewLabel("Weapons"), weaponSettings, weaponDetails),
	//		widget.NewLabel("Berths and Crew"), widget.NewHBox(berthSettings, berthDetails, vehicleSettings, vehicleDetails),
	//ui := widget.NewHBox(widget.NewVBox(widget.NewLabel("Drives"), shipSettings, widget.NewLabel("Weapons"), weaponSettings,
	//	widget.NewLabel("Berths"), berthSettings),
	//	widget.NewVBox(widget.NewLabel("Vehicles"), vehicleSettings),
	//	widget.NewVBox(shipDetails, weaponDetails, berthDetails, vehicleDetails))

	ui := widget.NewHBox(
		widget.NewVBox(widget.NewLabel("Drives"), shipSettings, widget.NewLabel("Berths"), berthSettings, widget.NewLabel("Weapons"), weaponSettings),
		widget.NewVBox(widget.NewLabel("Vehicles"), vehicleSettings),
		widget.NewVBox(shipDetails, weaponDetails, berthDetails, vehicleDetails))

	w.SetContent(ui)

	w.ShowAndRun()
	ui = widget.NewHBox(widget.NewVBox(widget.NewLabel("Drives"), widget.NewLabel("Weapons"),
		widget.NewLabel("Berths"), berthSettings),
		widget.NewVBox(widget.NewLabel("Vehicles"), vehicleSettings),
		widget.NewVBox(shipDetails, weaponDetails, berthDetails, vehicleDetails))

}

func saveMe() {

}

func loadMe() {

}
