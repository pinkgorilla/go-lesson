package main

import "log"

type Weekday int

const (
	Sunday Weekday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (w Weekday) String() string {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	return days[w-1]
}

type Size int

const (
	Small Size = iota + 1
	Reguler
	Large
)

func DataTypeIota() {
	log.Println("Data Type - iota")
	log.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday)
	log.Println(Small, Reguler, Large)
}
