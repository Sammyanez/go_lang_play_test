package main

import "os"

type Character interface {
	Name() string
	SetName(name string)
	Race() string
	SetRace(race string)
	Level() int
	SetLevel(level int)
	Health() int
	SetHealth(health int)
	calculateAttack() int
	ShowInfo()
	setCapacity(int)
	ExtractInfo(f *os.File)
}
