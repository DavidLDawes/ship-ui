package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func buildDetails() {
	buildShip()
	buildDrives()
	BuildTotal()
}

func main() {

	buildDetails()

	a := app.New()
	w := a.NewWindow("Designer")

	tlSelect = widget.NewSelect(TechLevels, techLevelChanged)
	tlSelect.SetSelected("F")

	tonsSelect = widget.NewSelect(Tons, tonsChanged)
	tonsSelect.SetSelected("200")

	jumpSelect = widget.NewSelect(EngineLevel, jumpChanged)
	jumpSelect.SetSelected("2")

	maneuverSelect = widget.NewSelect(EngineLevel, maneuverChanged)
	maneuverSelect.SetSelected("2")

	powerSelect = widget.NewSelect(EngineLevel, powerChanged)
	powerSelect.SetSelected("2")

	missileSelect = widget.NewSelect(EngineLevel, missileChanged)
	missileSelect.SetSelected("0")
	beamSelect = widget.NewSelect(EngineLevel, beamChanged)
	beamSelect.SetSelected("0")
	pulseSelect = widget.NewSelect(EngineLevel, pulseChanged)
	pulseSelect.SetSelected("0")
	fusionSelect = widget.NewSelect(EngineLevel, fusionChanged)
	fusionSelect.SetSelected("0")
	plasmaSelect = widget.NewSelect(EngineLevel, plasmaChanged)
	plasmaSelect.SetSelected("0")
	sandSelect = widget.NewSelect(EngineLevel, sandChanged)
	sandSelect.SetSelected("0")
	particleSelect = widget.NewSelect(EngineLevel, particleChanged)
	particleSelect.SetSelected("0")

	shipSettings := widget.NewForm(
		widget.NewFormItem("Tech Level", tlSelect),
		widget.NewFormItem("Tons", tonsSelect),
		widget.NewFormItem("Jump", jumpSelect),
		widget.NewFormItem("Maneuver", maneuverSelect),
		widget.NewFormItem("Power", powerSelect),
	)

	shipDetails := widget.NewVBox(widget.NewLabel("Ship Details"),
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

	weaponSettings := widget.NewForm(
		widget.NewFormItem("Missile", missileSelect),
		widget.NewFormItem("Beam", beamSelect),
		widget.NewFormItem("Pulse", pulseSelect),
		widget.NewFormItem("Fusion", fusionSelect),
		widget.NewFormItem("Sand", sandSelect),
		widget.NewFormItem("Plasma", plasmaSelect),
		widget.NewFormItem("Accelerators", particleSelect),
	)

	weaponDetails := widget.NewVBox(widget.NewLabel("Weapon Details"),
		&detailMissile,
		&detailBeam,
		&detailPulse,
		&detailFusion,
		&detailSand,
		&detailPlasma,
		&detailParticle,
	)

	ui := widget.NewHBox(widget.NewLabel("Designer"),
		shipSettings,
		shipDetails,
		weaponSettings,
		weaponDetails,
	)
	w.SetContent(ui)

	w.ShowAndRun()
}
