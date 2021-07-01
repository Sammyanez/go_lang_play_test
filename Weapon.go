package main

import "fmt"

type Weapon struct {
	name                string
	damage, staminaCost int
}

func (w *Weapon) Name() string {
	return w.name
}

func (w *Weapon) SetName(name string) {
	w.name = name
}

func (w *Weapon) Damage() int {
	return w.damage
}

func (w *Weapon) SetDamage(damage int) {
	w.damage = damage
}

func (w *Weapon) StaminaCost() int {
	return w.staminaCost
}

func (w *Weapon) SetStaminaCost(staminaCost int) {
	w.staminaCost = staminaCost
}

func (w *Weapon) Equip(name string, staminaCost int, damage int) {
	w.staminaCost = staminaCost
	w.name = name
	w.damage = damage
}

func (w *Weapon) printName() {
	fmt.Println("Weapon: ", w.name, "\n")
}
