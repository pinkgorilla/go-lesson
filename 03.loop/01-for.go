package main

import (
	"log"
	"time"
)

func ForBasic() {
	log.Println("For - basic")
	for i := 0; i < 3; i++ {
		log.Println(i)
	}
}

func ForRange() {
	log.Println("For - range")
	arr := []string{"a", "b", "c"}

	// for index := 0; index < len(arr); index++ {
	// 	item := arr[index]
	// 	log.Printf("index: %d,item: %v\n", index, item)
	// }

	for index, item := range arr {
		log.Printf("index: %d,item: %v\n", index, item)
	}
}

func ForEver() {
	after := time.After(2 * time.Second)
	log.Println("For - Ever")
	for {
		select {
		case <-after:
			log.Println("timeout!")
			return
		default:
			time.Sleep(time.Millisecond * 100)
			log.Println("loop...")
		}
	}
}
