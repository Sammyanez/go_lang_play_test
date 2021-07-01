package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Warrior struct {
	Character
	name, race    string
	level, health int
	myWeapon      Weapon
	stamina       int
}

func (w *Warrior) Name() string {
	return w.name
}

func (w *Warrior) SetName(name string) {
	w.name = name
}

func (w *Warrior) Race() string {
	return w.race
}

func (w *Warrior) SetRace(race string) {
	w.race = race
}

func (w *Warrior) Level() int {
	return w.level
}

func (w *Warrior) SetLevel(level int) {
	w.level = level
}

func (w *Warrior) Health() int {
	return w.health
}

func (w *Warrior) SetHealth(health int) {
	w.health = health
}

func (w *Warrior) Stamina() int {
	return w.stamina
}

func (w *Warrior) setCapacity(stamina int) {
	w.stamina = stamina
}

func (w *Warrior) calculateAttack() int {

	if w.myWeapon.staminaCost > w.stamina {
		return 0
	} else {
		println(w.name, " attacks with their", w.myWeapon.name)
		println("It does ", w.myWeapon.damage, " damage!")
		w.setCapacity(w.stamina - w.myWeapon.staminaCost)
		return w.myWeapon.damage
	}

}

func (w *Warrior) ShowInfo() {
	fmt.Println("Character Stats")
	fmt.Println("Name:", w.Name())
	fmt.Println("Race:", w.Race())
	fmt.Println("Level:", w.Level())
	fmt.Println("Health:", w.Health())
	fmt.Println("Class: Warrior")
	fmt.Println("Stamina:", w.stamina)
	w.myWeapon.printName()

}

func (w *Warrior) ExtractInfo(f *os.File) {
	_, err2 := f.WriteString("\n")
	_, err2 = f.WriteString("Character Stats")
	_, err2 = f.WriteString("\n")
	if err2 != nil {
		log.Fatal(err2)
	}
	_, err2 = f.WriteString("Name:")
	_, err2 = f.WriteString(w.Name())
	_, err2 = f.WriteString("\n")
	_, err2 = f.WriteString("Race:")
	_, err2 = f.WriteString(w.Race())
	_, err2 = f.WriteString("\n")
	_, err2 = f.WriteString("Level:")
	_, err2 = f.WriteString(strconv.Itoa(w.Level()))
	_, err2 = f.WriteString("\n")
	_, err2 = f.WriteString("Health:")
	_, err2 = f.WriteString(strconv.Itoa(w.Health()))
	_, err2 = f.WriteString("\n")
	_, err2 = f.WriteString("Class: Witch")
	_, err2 = f.WriteString("\n")
	_, err2 = f.WriteString("Stamina:")
	_, err2 = f.WriteString(strconv.Itoa(w.stamina))
	_, err2 = f.WriteString("\n")
	_, err2 = f.WriteString("Weapon: ")
	_, err2 = f.WriteString(w.myWeapon.name)
	_, err2 = f.WriteString("\n")
}

func (w *Warrior) equipWeapon(we *Weapon) {

	w.myWeapon.Equip(we.name, we.staminaCost, we.damage)

}
