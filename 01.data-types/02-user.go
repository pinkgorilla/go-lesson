package main

import (
	"fmt"
	"log"
)

// Define Person Type
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("\nPerson:\n--------\n Name:%s\n Age:%d", p.Name, p.Age)
}

// Define Man type of Person type
type Man Person

func (m Man) String() string {
	return fmt.Sprintf("\nMan:\n--------\n Name:%s\n Age:%d", m.Name, m.Age)
}

// Define Developer Type
// Promoted Fields / type embedding
type Developer struct {
	Person
	Language string
}

func (d Developer) String() string {
	return fmt.Sprintf("\nDeveloper:\n--------\n Name:%s\n Age:%d\n Language:%s", d.Name, d.Age, d.Language)
}

func DataTypeComplex() {
	log.Println("Data Type - complex")
	var p Person
	var m Man
	var d Developer
	p = Person{
		Name: "John Doe",
		Age:  24,
	}
	log.Println(p)

	m = Man{"Adam", 1}
	m = Man{
		Name: "Adam",
		Age:  1,
	}
	log.Println(m)

	d = Developer{
		Person: Person{
			Name: "Neo",
			Age:  25,
		},
		Language: "Go",
	}
	log.Println(d)
}
