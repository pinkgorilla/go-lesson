package main

import "log"

func DataTypePrimitives() {
	log.Println("Data Type - primitives")
	var i int
	var s string
	var f float64

	i = 10
	s = "let's \"Go\"!"
	s = `let's "Go"!`
	f = .0001

	log.Println(i, s, f)
}

func DataTypeZeroValues() {
	log.Println("Data Type - zero values")
	var i int
	var pi *int
	log.Printf("int:%v", i)
	log.Printf("*int:%v", pi)

	var s string
	var ps *string
	log.Printf("string:%v", s)
	log.Printf("*string:%v", ps)

	var p Person
	var pp *Person
	log.Printf("person:%v", p)
	log.Printf("*person:%v", pp)

	var f func()
	var pf *func()
	log.Printf("func:%v", f)
	log.Printf("*func:%v", pf)

}
