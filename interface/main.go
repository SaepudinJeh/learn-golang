package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("hello", hasName.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string { //Harus return string sesuai dengan kontrak
	return person.name
}

func main() {
	saepudin := Person{
		name: "Agung",
	}

	sayHello(saepudin)
}
