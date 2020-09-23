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
	vehiclesInit()

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
		widget.NewVBox(widget.NewLabel("Drives"), shipSettings, widget.NewLabel("Berths & Crew"), berthSettings, widget.NewLabel("Weapons"), weaponSettings),
		widget.NewVBox(widget.NewLabel("Vehicles"), vehicleSettings),
		widget.NewVBox(shipDetails, weaponDetails, berthDetails, vehicleDetails))

	w.SetContent(ui)

	w.ShowAndRun()
}

func saveMe() {

}

func loadMe() {

}
