package main

import "log"

func IfBasic() {
	log.Println("If - basic")
	b := true
	if b {
		log.Println("hello")
	}
}

func IfInline() {
	log.Println("If - inline")
	fn := func() (int, error) {
		return 0, nil
	}
	if i, e := fn(); i < 1 || e != nil {
		log.Println("zero or error")
	}

	i, e := fn()
	if  i < 1 || e != nil {
		log.Println("zero or error")
	}	
}
