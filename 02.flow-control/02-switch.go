package main

import "log"

func SwitchBasic() {
	log.Println("Switch - basic")
	v := 1
	switch v {
	case 1:
		log.Println("1")
	case 2:
		log.Println("2")
	case 3:
		log.Println("3")
	default:
		log.Println("default")
	}
}

func SwitchExpression() {
	log.Println("Switch - expression")
	v := 2
	switch {
	case v%2 == 0:
		log.Println("even")
	default:
		log.Println("odd")
	}
}

func SwitchFallthrough() {
	log.Println("Switch - fallthrough")
	v := 3
	switch v {
	case 1:
		log.Println("1")
		fallthrough
	case 2:
		log.Println("2")
		fallthrough
	case 3:
		log.Println("3")
		fallthrough
	case 4:
		log.Println("4")
		fallthrough
	case 5:
		log.Println("5")
	default:
		log.Println("too big")
	}
}
