package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	game := Game{}
	myWarrior := Warrior{}
	myWitch := Witch{}
	myWeapon := Weapon{}
	mySpell := Spell{}
	chSend := make(chan string)
	var myFile string

	println("What is your file name: ")
	fmt.Scanln(&myFile)

	wg.Add(2)
	go func(chSend chan<- string) {
		defer wg.Done()
		f, err := os.Open(myFile)

		if err != nil {
			log.Fatal(err)
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {

			chSend <- scanner.Text()

		}

		close(chSend)

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}(chSend)

	go func(chSend <-chan string) {
		var i string
		i = <-chSend
		if i == "Warrior" {

			myWarrior.SetName(<-chSend)
			myWarrior.SetRace(<-chSend)
			temp, _ := strconv.Atoi(<-chSend)
			myWarrior.SetLevel(temp)
			temp, _ = strconv.Atoi(<-chSend)
			myWarrior.SetHealth(temp)
			temp, _ = strconv.Atoi(<-chSend)
			myWarrior.setCapacity(50)
			myWeapon.SetName(<-chSend)
			temp, _ = strconv.Atoi(<-chSend)
			myWeapon.SetDamage(temp)
			temp, _ = strconv.Atoi(<-chSend)
			myWeapon.SetStaminaCost(temp)
			myWarrior.equipWeapon(&myWeapon)
		}
		i = <-chSend
		i = <-chSend
		if i == "Witch" {

			myWitch.SetName(<-chSend)
			myWitch.SetRace(<-chSend)
			temp, _ := strconv.Atoi(<-chSend)
			myWitch.SetLevel(temp)
			temp, _ = strconv.Atoi(<-chSend)
			myWitch.SetHealth(temp)
			temp, _ = strconv.Atoi(<-chSend)
			myWitch.setCapacity(50)
			temp, _ = strconv.Atoi(<-chSend)
			myWitch.SetNumSpells(temp)
			for j := 0; j < myWitch.numSpells; j++ {
				mySpell.SetSpellName(<-chSend)
				temp, _ = strconv.Atoi(<-chSend)
				mySpell.SetDamage(temp)
				temp, _ = strconv.Atoi(<-chSend)
				mySpell.SetManaCost(temp)
				myWitch.addSpell(mySpell)
			}
			wg.Done()

		}

	}(chSend)
	var character1 Character = &myWarrior
	var character2 Character = &myWitch
	game.AddCharacter(character1)
	game.AddCharacter(character2)

	Menu(&game)

}
