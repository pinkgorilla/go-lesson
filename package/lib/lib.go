package lib

import "fmt"

// Define Person Type
type Person struct {
	Name    string
	Age     int
	Address string
}

func (this Person) Say() {
	this.private()
}
func (this Person) private() {
	fmt.Println(this.Name)
}
func (this *Person) ChangeName(newName string) {
	this.Name = newName
}

// func (p Person) String() string {
// 	return fmt.Sprintf("\nPerson:\n--------\n Name:%s\n Age:%d", p.Name, p.Age)
// }
func Say(message string) {
	fmt.Println(message)
}

func private() {

}

type Beo struct {
}

func (b Beo) Say() {
	fmt.Println("beo")
}

type X struct {
}
