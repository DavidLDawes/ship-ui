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

var Weapons = weaponDetails{
	missile:     0,
	beam:        0,
	pulse:       0,
	plasma:      0,
	fusion:      0,
	sandcaster:  0,
	accelerator: 0,
}

type Weapon struct {
	name     string
	max      int
	tons     float32
	ammoName string
	ammoTons int
}

var missle = Weapon{
	name:     "Missile launcher turret",
	max:      3,
	tons:     1.0,
	ammoName: "missiles",
	ammoTons: 4,
}
var beam = Weapon{
	name:     "Beam laser turret",
	max:      3,
	tons:     1.0,
	ammoName: "",
	ammoTons: 0,
}
var pulse = Weapon{
	name:     "Pulse laser turret",
	max:      3,
	tons:     1.0,
	ammoName: "",
	ammoTons: 0,
}
var fusion = Weapon{
	name:     "Fusion gun turret",
	max:      2,
	tons:     2.0,
	ammoName: "",
	ammoTons: 0,
}
var sand = Weapon{
	name:     "Sandcaster turret",
	max:      3,
	tons:     .5,
	ammoName: "Sand",
	ammoTons: 1,
}
var plasma = Weapon{
	name:     "Plasma gun turret",
	max:      2,
	tons:     1.5,
	ammoName: "",
	ammoTons: 0,
}
var particle = Weapon{
	name:     "Particle accelerator turret",
	max:      1,
	tons:     3.0,
	ammoName: "",
	ammoTons: 0,
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

// JumpChanged handle updates to the jump drive
func missileChanged(value string) {
	missiles, err := strconv.Atoi(value)
	if err == nil {
		Weapons.missile = missiles
	}
	buildMissile()
	BuildTotal()
}

func beamChanged(value string) {
	beamTurrets, err := strconv.Atoi(value)
	if err == nil {
		Weapons.beam = beamTurrets
	}
	buildBeam()
	BuildTotal()
}

func pulseChanged(value string) {
	pulse, err := strconv.Atoi(value)
	if err == nil {
		Weapons.pulse = pulse
	}
	buildPulse()
	BuildTotal()
}

func fusionChanged(value string) {
	fusion, err := strconv.Atoi(value)
	if err == nil {
		Weapons.fusion = fusion
	}
	buildFusion()
	BuildTotal()
}

func sandChanged(value string) {
	sand, err := strconv.Atoi(value)
	if err == nil {
		Weapons.sandcaster = sand
	}
	buildSand()
	BuildTotal()
}

func plasmaChanged(value string) {
	plasma, err := strconv.Atoi(value)
	if err == nil {
		Weapons.plasma = plasma
	}
	buildPlasma()
	BuildTotal()
}

func particleChanged(value string) {
	particle, err := strconv.Atoi(value)
	if err == nil {
		Weapons.plasma = particle
	}
	buildParticle()
	BuildTotal()
}

func buildMissile() {
	detailMissile.SetText(fmt.Sprintf("Triple missile turrets: %d, tons: %2.1f, ammo tons: %d", Weapons.missile, float32(Weapons.missile), 4.0*float32(Weapons.missile)))
	detailMissile.Refresh()
}
func buildBeam() {
	detailBeam.SetText(fmt.Sprintf("Triple beam laser turrets: %d, tons: %2.1f", Weapons.beam, float32(Weapons.beam)))
	detailBeam.Refresh()
}
func buildPulse() {
	detailPulse.SetText(fmt.Sprintf("Triple pulse laser turrets: %d, tons: %2.1f", Weapons.pulse, float32(Weapons.pulse)))
	detailPulse.Refresh()
}
func buildPlasma() {
	detailPlasma.SetText(fmt.Sprintf("Double plasma gun turrets: %d, tons: %2.1f", Weapons.plasma, 2.0*float32(Weapons.plasma)))
	detailPulse.Refresh()
}
func buildFusion() {
	detailFusion.SetText(fmt.Sprintf("Double fusion gun turrets: %d, tons: %2.1f", Weapons.fusion, 2.0*float32(Weapons.fusion)))
	detailFusion.Refresh()
}
func buildSand() {
	if Weapons.sandcaster > 0 {
		detailSand.SetText(fmt.Sprintf("Triple sandcaster turrets: %d, tons: %2.1f, ammo tons %2.1f", Weapons.sandcaster, 0.5*float32(Weapons.sandcaster), 1.0*float32(Weapons.sandcaster)))
	} else {
		detailSand.SetText("Triple sandcaster turrets: 0, tons: 0")
	}
	detailFusion.Refresh()
}
func buildParticle() {
	detailParticle.SetText(fmt.Sprintf("Paticle acceleraor turrets: %d, tons: %2.1f", Weapons.accelerator, 3.0*float32(Weapons.accelerator)))
	detailFusion.Refresh()

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
	return int(5.0*float32(Weapons.missile) + float32(Weapons.beam) + float32(Weapons.pulse) + 2.0*float32(Weapons.fusion) + 2.0*float32(Weapons.plasma) + 5.0*float32(Weapons.accelerator))
}
