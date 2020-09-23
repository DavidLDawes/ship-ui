package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

type weaponCounts struct {
	missile     int
	beam        int
	pulse       int
	plasma      int
	fusion      int
	sandcaster  int
	accelerator int
}

var weapons = weaponCounts{
	missile:     0,
	beam:        0,
	pulse:       0,
	plasma:      0,
	fusion:      0,
	sandcaster:  0,
	accelerator: 0,
}

var weaponSettings *widget.Form

var detailMissile *widget.Label = widget.NewLabel("")
var detailBeam *widget.Label = widget.NewLabel("")
var detailPulse *widget.Label = widget.NewLabel("")
var detailFusion *widget.Label = widget.NewLabel("")
var detailSand *widget.Label = widget.NewLabel("")
var detailPlasma *widget.Label = widget.NewLabel("")
var detailParticle *widget.Label = widget.NewLabel("")

var weaponDetails *widget.Box = widget.NewVBox()

var missileSelect *widget.Select
var beamSelect *widget.Select
var pulseSelect *widget.Select
var fusionSelect *widget.Select
var sandSelect *widget.Select
var plasmaSelect *widget.Select
var particleSelect *widget.Select

var weaponsSelect []*widget.Select

var ignoreMissile = false
var ignoreBeam = false
var ignorePulse = false
var ignorePlasma = false
var ignoreSand = false
var ignoreFusion = false
var ignoreParticle = false

func weaponsInit() {
	missileSelect = widget.NewSelect(weaponLevel, nothing)
	missileSelect.SetSelected("0")
	missileSelect.OnChanged = missileChanged

	beamSelect = widget.NewSelect(weaponLevel, nothing)
	beamSelect.SetSelected("0")
	missileSelect.OnChanged = beamChanged

	pulseSelect = widget.NewSelect(weaponLevel, nothing)
	pulseSelect.SetSelected("0")
	missileSelect.OnChanged = pulseChanged

	fusionSelect = widget.NewSelect(weaponLevel, nothing)
	fusionSelect.SetSelected("0")
	missileSelect.OnChanged = fusionChanged

	plasmaSelect = widget.NewSelect(weaponLevel, nothing)
	plasmaSelect.SetSelected("0")
	missileSelect.OnChanged = plasmaChanged

	sandSelect = widget.NewSelect(weaponLevel, nothing)
	sandSelect.SetSelected("0")
	missileSelect.OnChanged = sandChanged

	particleSelect = widget.NewSelect(weaponLevel, nothing)
	particleSelect.SetSelected("0")
	missileSelect.OnChanged = particleChanged

	weaponSettings = widget.NewForm(
		widget.NewFormItem("Missile", missileSelect),
		widget.NewFormItem("Beam", beamSelect),
		widget.NewFormItem("Pulse", pulseSelect),
		widget.NewFormItem("Fusion", fusionSelect),
		widget.NewFormItem("Sand", sandSelect),
		widget.NewFormItem("Plasma", plasmaSelect),
		widget.NewFormItem("Accelerators", particleSelect),
	)
}

func weaponsSelectInit() {
	weaponsSelect = make([]*widget.Select, 7)
	weaponsSelect[0] = missileSelect
	weaponsSelect[1] = beamSelect
	weaponsSelect[2] = pulseSelect
	weaponsSelect[3] = fusionSelect
	weaponsSelect[4] = sandSelect
	weaponsSelect[5] = plasmaSelect
	weaponsSelect[6] = particleSelect

	setWeaponDetails()
}

func missileChanged(value string) {
	if !ignoreMissile {
		missiles, err := strconv.Atoi(value)
		if err == nil {
			weapons.missile = missiles
			if countWeapons() > StarShip.hardpoints {
				weapons.missile = missiles - countWeapons() + StarShip.hardpoints
				if weapons.missile < 0 {
					weapons.missile = 0
				}
				if weapons.missile != missiles {
					missileSelect.SetSelected(strconv.Itoa(weapons.missile))
				}
			}
			setWeaponDetails()
		}
		buildMissile()
		buildCrew()
		buildTotal()
	}
}

func beamChanged(value string) {
	if !ignoreBeam {
		beamTurrets, err := strconv.Atoi(value)
		if err == nil {
			weapons.beam = beamTurrets
			if countWeapons() > StarShip.hardpoints {
				weapons.beam = beamTurrets - countWeapons() + StarShip.hardpoints
				if weapons.beam < 0 {
					weapons.beam = 0
				}
				beamSelect.SetSelected(strconv.Itoa(weapons.beam))
			}
			setWeaponDetails()
		}
		buildBeam()
		buildCrew()
		buildTotal()
	}
}

func pulseChanged(value string) {
	if !ignorePulse {
		pulse, err := strconv.Atoi(value)
		if err == nil {
			weapons.pulse = pulse
			if countWeapons() > StarShip.hardpoints {
				weapons.pulse = pulse - countWeapons() + StarShip.hardpoints
				if weapons.pulse < 0 {
					weapons.pulse = 0
				}
				pulseSelect.SetSelected(strconv.Itoa(weapons.pulse))
			}
			setWeaponDetails()
		}
		buildPulse()
		buildCrew()
		buildTotal()
	}
}

func fusionChanged(value string) {
	if !ignoreFusion {
		fusion, err := strconv.Atoi(value)
		if err == nil {
			weapons.fusion = fusion
			if countWeapons() > StarShip.hardpoints {
				weapons.fusion = fusion - countWeapons() + StarShip.hardpoints
				if weapons.fusion < 0 {
					weapons.fusion = 0
				}
				fusionSelect.SetSelected(strconv.Itoa(weapons.fusion))
			}
			setWeaponDetails()
		}
		buildFusion()
		buildCrew()
		buildTotal()
	}
}

func sandChanged(value string) {
	if !ignoreSand {
		sand, err := strconv.Atoi(value)
		if err == nil {
			weapons.sandcaster = sand
			if countWeapons() > StarShip.hardpoints {
				weapons.sandcaster = sand - countWeapons() + StarShip.hardpoints
				if weapons.sandcaster < 0 {
					weapons.sandcaster = 0
				}
				sandSelect.SetSelected(strconv.Itoa(weapons.sandcaster))
			}
			setWeaponDetails()
		}
		buildSand()
		buildCrew()
		buildTotal()
	}
}

func plasmaChanged(value string) {
	if !ignorePlasma {
		plasma, err := strconv.Atoi(value)
		if err == nil {
			weapons.plasma = plasma
			if countWeapons() > StarShip.hardpoints {
				weapons.plasma = plasma - countWeapons() + StarShip.hardpoints
				if weapons.plasma < 0 {
					weapons.plasma = 0
				}
				plasmaSelect.SetSelected(strconv.Itoa(weapons.plasma))
			}
			setWeaponDetails()
		}
		buildPlasma()
		buildCrew()
		buildTotal()
	}
}

func particleChanged(value string) {
	if !ignoreParticle {
		particle, err := strconv.Atoi(value)
		if err == nil {
			weapons.accelerator = particle
			if countWeapons() > StarShip.hardpoints {
				weapons.accelerator = particle - countWeapons() + StarShip.hardpoints
				if weapons.accelerator < 0 {
					weapons.accelerator = 0
				}
				particleSelect.SetSelected(strconv.Itoa(weapons.accelerator))
			}
		}
		buildParticle()
		buildCrew()
		buildTotal()
	}
}

func buildMissile() {
	detailMissile.SetText(buildAmmoWeaponString("Triple Missile turrets: %d, tons: %d, ammo tons: %d", weapons.missile, int(armor()*float32(weapons.missile)+.9999), int(armor()*float32(4*weapons.missile)+.9999)))
	detailMissile.Refresh()
}
func buildBeam() {
	detailBeam.SetText(buildWeaponString("Triple Beam laser turrets: %d, tons: %d", weapons.beam, int(armor()*float32(weapons.beam)+.9999)))
	detailBeam.Refresh()
}
func buildPulse() {
	detailPulse.SetText(buildWeaponString("Triple Pulse lasr turrets: %d, tons: %d", weapons.pulse, int(armor()*float32(weapons.pulse)+.9999)))
	detailPulse.Refresh()
}
func buildPlasma() {
	detailPlasma.SetText(buildWeaponString("Double Plasma gun turrets: %d, tons: %d", weapons.plasma, int(armor()*float32(2*weapons.plasma)+.9999)))
	detailPulse.Refresh()
}
func buildFusion() {
	detailFusion.SetText(buildWeaponString("Double Fusion gun turrets: %d, tons: %d", weapons.fusion, int(armor()*float32(2*weapons.fusion)+.9999)))
	detailFusion.Refresh()
}
func buildSand() {
	detailSand.SetText(buildAmmoWeaponString("Triple Sandcaster turrets: %d, tons: %d, ammo tons: %d", weapons.sandcaster, int(armor()*float32(weapons.sandcaster)/2.0+.9999), int(armor()*float32(weapons.sandcaster)+.9999)))
	detailFusion.Refresh()
}
func buildParticle() {
	detailParticle.SetText(buildWeaponString("Particle Accelerator turrets: %d, tons: %d", weapons.accelerator, int(armor()*float32(3*weapons.accelerator)+.9999)))
	detailFusion.Refresh()
}

func countWeapons() int {
	result := weapons.missile + weapons.beam + weapons.pulse + weapons.plasma + weapons.sandcaster + weapons.fusion + weapons.accelerator
	return result
}

func buildWeapons() {
	buildMissile()
	buildBeam()
	buildPlasma()
	buildFusion()
	buildSand()
	buildParticle()
}

func weaponsTonsUsed() int {
	result := int(.9999 + armor()*(5.0*float32(weapons.missile)+float32(weapons.beam)+float32(weapons.pulse)+2.0*float32(weapons.fusion)+2.0*float32(weapons.plasma)+5.0*float32(weapons.accelerator)))
	return result
}

func buildWeaponString(weaponDescription string, count int, tons int) string {
	if count > 0 {
		return fmt.Sprintf(weaponDescription, count, int(float32(tons)*armor()))
	}
	return ""
}

func buildAmmoWeaponString(weaponAmmoDescription string, count int, tons int, ammoTons int) string {
	if count > 0 {
		return fmt.Sprintf(weaponAmmoDescription, count, int(.999+float32(tons)*armor()), int(.999+float32(ammoTons)*armor()))
	}
	return ""
}

func setWeaponDetails() {
	// Start again
	weaponDetails.Children = make([]fyne.CanvasObject, 0)
	// Check each and add if needed
	addWeapon(weapons.missile, weaponDetails, detailMissile)
	addWeapon(weapons.beam, weaponDetails, detailBeam)
	addWeapon(weapons.pulse, weaponDetails, detailPulse)
	addWeapon(weapons.plasma, weaponDetails, detailPlasma)
	addWeapon(weapons.fusion, weaponDetails, detailFusion)
	addWeapon(weapons.sandcaster, weaponDetails, detailSand)
	addWeapon(weapons.accelerator, weaponDetails, detailParticle)
	weaponDetails.Refresh()
}

// Add next weapon, if needed, to the detailed list of weapons
func addWeapon(count int, box *widget.Box, label *widget.Label) {
	if count > 0 {
		weaponDetails.Children = append(weaponDetails.Children, label)
	}
}

func armor() float32 {
	if StarShip.armored {
		return 1.1
	}
	return 1.0
}
