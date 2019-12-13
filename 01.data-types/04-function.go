package main

import "log"

type Actor struct {
	Act func()
}

func (a Actor) Action() {
	a.Act()
}

func DataTypeFunction() {
	log.Println("Data Type - function")
	john := Actor{}
	cook := func() {
		log.Println("cooking!!")
	}
	box := func() {
		log.Println("boxing!!")
	}

	john.Act = cook
	john.Action()

	john.Act = box
	john.Action()
}
