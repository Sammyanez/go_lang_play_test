package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func Menu(g *Game) {

	choice := 0
	string0 := "\nPlease choose an option"
	string1 := "1 - FIGHT!"
	string2 := "2 - Print All Characters"
	string3 := "3 - Exit"

	f, err := os.Create("data.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err2 := f.WriteString("Recap\n")

	if err2 != nil {
		log.Fatal(err2)
	}

Loop:
	for {

		fmt.Println(string0)
		fmt.Println(string1)
		fmt.Println(string2)
		fmt.Println(string3)

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			g.Attack()
			g.SetCurrentTurn(g.currentTurn + 1)
			_, err2 := f.WriteString("\n")
			_, err2 = f.WriteString("Round: ")
			_, err2 = f.WriteString(strconv.Itoa(g.currentTurn))
			_, err2 = f.WriteString("\n")
			if err2 != nil {
				log.Fatal(err2)
			}
			g.writeCharacters(f)
		case 2:
			g.printCharacters()
		case 3:
			break Loop

		}
	}
}
