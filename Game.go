package main

import (
	"os"
)

type Game struct {
	Character
	Witch
	Warrior
	players                    []Character
	currentTurn, numCharacters int
}

func (g *Game) CurrentTurn() int {
	return g.currentTurn
}

func (g *Game) SetCurrentTurn(currentTurn int) {
	g.currentTurn = currentTurn
}

func (g *Game) NumCharacters() int {
	return g.numCharacters
}

func (g *Game) SetNumCharacters(numCharacters int) {
	g.numCharacters = numCharacters
}

func (g *Game) AddCharacter(c Character) {
	g.players = append(g.players, c)
	g.numCharacters++
}

func (g *Game) printCharacters() {

Loop:
	for i := range g.players {
		if g.players[i] == nil {
			break Loop
		}
		g.players[i].ShowInfo()

	}
}

func (g *Game) writeCharacters(f *os.File) {

Loop:
	for i := range g.players {
		if g.players[i] == nil {
			break Loop
		}
		g.players[i].ExtractInfo(f)

	}
}

func (g *Game) Attack() {

	if g.checkMortality() == 1 || g.checkMortality() == 2 {
		return
	} else {
		var p1 int = g.players[0].calculateAttack()
		var p2 int = g.players[1].calculateAttack()

		if p1 <= 0 {
			println("Oh, No! ", g.players[0].Name(), " couldn't attack ")
		} else if p2 <= 0 {
			println("Oh, No! ", g.players[1].Name(), " couldn't attack ")
		} else {
		}

		g.players[0].SetHealth(g.players[0].Health() - p2)
		g.players[1].SetHealth(g.players[1].Health() - p1)

		a := g.checkMortality()

		if a-1 == 0 {
			println(g.players[1].Name(), " has Won \n They are awarded with a Full Recovery Potion")
			g.players[1].SetHealth(100)
			g.players[1].setCapacity(100)
		} else if a-1 == 1 {
			println(g.players[0].Name(), " has Won \n They are awarded with a Full Recovery Potion")
			g.players[0].SetHealth(100)
			g.players[0].setCapacity(100)

		}

	}
}

func (g *Game) checkMortality() int {
	if g.players[0].Health() <= 0 {
		println(g.players[0].Name(), " is Dead!")
		println("Maybe it is time to replace them...")
		return 1
	}

	if g.players[1].Health() <= 0 {
		println(g.players[1].Name(), " is Dead! \n Maybe it is time to replace them...")
		return 2
	}

	return 0
}
