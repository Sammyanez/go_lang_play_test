package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type Witch struct {
	Character
	name, race      string
	level, health   int
	Spells          []Spell
	numSpells, mana int
}

func (w *Witch) Name() string {
	return w.name
}

func (w *Witch) SetName(name string) {
	w.name = name
}

func (w *Witch) Race() string {
	return w.race
}

func (w *Witch) SetRace(race string) {
	w.race = race
}

func (w *Witch) Level() int {
	return w.level
}

func (w *Witch) SetLevel(level int) {
	w.level = level
}

func (w *Witch) Health() int {
	return w.health
}

func (w *Witch) SetHealth(health int) {
	w.health = health
}

func (w *Witch) NumSpells() int {
	return w.numSpells
}

func (w *Witch) SetNumSpells(numSpells int) {
	w.numSpells = numSpells
}

func (w *Witch) Mana() int {
	return w.mana
}

func (w *Witch) setCapacity(mana int) {
	w.mana = mana
}

func (w *Witch) calculateAttack() int {

	var choice = rand.Intn(len(w.Spells))

	if w.Spells[choice].manaCost >= w.mana {
		return 0
	} else {
		println(w.name, " casts ", w.Spells[choice].spellName)
		println("It does ", w.Spells[choice].damage, " damage!")
		w.setCapacity(w.mana - w.Spells[choice].manaCost)
		return w.Spells[choice].damage
	}

}

func (w *Witch) addSpell(s Spell) {
	w.Spells = append(w.Spells, s)

}

func (w *Witch) ShowInfo() {

	fmt.Println("Character Stats")
	fmt.Println("Name:", w.Name())
	fmt.Println("Race:", w.Race())
	fmt.Println("Level:", w.Level())
	fmt.Println("Health:", w.Health())
	fmt.Println("Class: Witch")
	fmt.Println("Mana:", w.mana)
	fmt.Println("Spells:")

	for i := range w.Spells {

		fmt.Println(i+1, ". ", w.Spells[i].spellName)
	}

}

func (w *Witch) ExtractInfo(f *os.File) {
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
	_, err2 = f.WriteString("Mana:")
	_, err2 = f.WriteString(strconv.Itoa(w.mana))
	_, err2 = f.WriteString("\n")
	_, err2 = f.WriteString("Spells:")
	_, err2 = f.WriteString("\n")

	for i := range w.Spells {

		_, err2 = f.WriteString(strconv.Itoa(i + 1))
		_, err2 = f.WriteString(". ")
		_, err2 = f.WriteString(w.Spells[i].spellName)
		_, err2 = f.WriteString("\n")
	}

}
