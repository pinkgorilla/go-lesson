package main

import (
	"fmt"

	l "../lib"
)

type Speakable interface {
	Say()
}

type Speakable2 interface {
	Say()
}

func main() {
	var i Speakable
	var j Speakable2
	i = l.Person{
		Name:    "Waris",
		Address: "Yado",
	}
	i = l.Beo{}
	j = i
	// concrete, ok:= inter.(type)
	// p, ok := i.(l.Person)
	// if ok {
	// 	p.Say()
	// }
	i.Say()
	j.Say()
	// p.Say()
	// p.ChangeName("Joni")
	// p.Say()
	fmt.Println(i)
	e()
}

func e() {
	var err error
	// err := fmt.Errorf("%s", "outer error")
	if true {
		err := fmt.Errorf("%s", "inner error")
		fmt.Println(err)
	}
	fmt.Println(err)
}
