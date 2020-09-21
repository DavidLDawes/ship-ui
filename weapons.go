package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/widget"
)

type weaponDetails struct {
	missile     int
	beam        int
	pulse       int
	plasma      int
	fusion      int
	sandcaster  int
	accelerator int
}

var weapons = weaponDetails{
	missile:     0,
	beam:        0,
	pulse:       0,
	plasma:      0,
	fusion:      0,
	sandcaster:  0,
	accelerator: 0,
}

var detailMissile widget.Label
var detailBeam widget.Label
var detailPulse widget.Label
var detailFusion widget.Label
var detailSand widget.Label
var detailPlasma widget.Label
var detailParticle widget.Label

var missileSelect *widget.Select
var beamSelect *widget.Select
var pulseSelect *widget.Select
var fusionSelect *widget.Select
var sandSelect *widget.Select
var plasmaSelect *widget.Select
var particleSelect *widget.Select

var weaponsSelect []*widget.Select

var weaponsAlreadyInit bool = false

var ignoreMissile = false
var ignoreBeam = false
var ignorePulse = false
var ignorePlasma = false
var ignoreSand = false
var ignoreFusion = false
var ignoreParticle = false

func weaponsInit() {
	if !weaponsAlreadyInit {
		weaponsAlreadyInit = true
		weaponsSelect = make([]*widget.Select, 7)
		weaponsSelect[0] = missileSelect
		weaponsSelect[1] = beamSelect
		weaponsSelect[2] = pulseSelect
		weaponsSelect[3] = fusionSelect
		weaponsSelect[4] = sandSelect
		weaponsSelect[5] = plasmaSelect
		weaponsSelect[6] = particleSelect
	}
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
	detailMissile.SetText(fmt.Sprintf("Triple missile turrets: %d, tons: %d, ammo tons: %d", weapons.missile, weapons.missile, 4*weapons.missile))
	detailMissile.Refresh()
}
func buildBeam() {
	detailBeam.SetText(fmt.Sprintf("Triple beam laser turrets: %d, tons: %2.1f", weapons.beam, float32(weapons.beam)))
	detailBeam.Refresh()
}
func buildPulse() {
	detailPulse.SetText(fmt.Sprintf("Triple pulse laser turrets: %d, tons: %2.1f", weapons.pulse, float32(weapons.pulse)))
	detailPulse.Refresh()
}
func buildPlasma() {
	detailPlasma.SetText(fmt.Sprintf("Double plasma gun turrets: %d, tons: %2.1f", weapons.plasma, 2.0*float32(weapons.plasma)))
	detailPulse.Refresh()
}
func buildFusion() {
	detailFusion.SetText(fmt.Sprintf("Double fusion gun turrets: %d, tons: %2.1f", weapons.fusion, 2.0*float32(weapons.fusion)))
	detailFusion.Refresh()
}
func buildSand() {
	if weapons.sandcaster > 0 {
		detailSand.SetText(fmt.Sprintf("Triple sandcaster turrets: %d, tons: %2.1f, ammo tons %2.1f", weapons.sandcaster, 0.5*float32(weapons.sandcaster), 1.0*float32(weapons.sandcaster)))
	} else {
		detailSand.SetText("Triple sandcaster turrets: 0, tons: 0")
	}
	detailFusion.Refresh()
}
func buildParticle() {
	detailParticle.SetText(fmt.Sprintf("Paticle acceleraor turrets: %d, tons: %2.1f", weapons.accelerator, 3.0*float32(weapons.accelerator)))
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
	result := int(5.0*float32(weapons.missile) + float32(weapons.beam) + float32(weapons.pulse) + 2.0*float32(weapons.fusion) + 2.0*float32(weapons.plasma) + 5.0*float32(weapons.accelerator))
	return result
}
