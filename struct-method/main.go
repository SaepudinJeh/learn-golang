package main

import "fmt"

type Customer struct {
	name, adress string
	age, number  int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "My name is", customer.name)
}

func main() {
	saepudin := Customer{
		name:   "Agung",
		adress: "Cirebon",
		age:    21,
		number: 32434,
	}

	saepudin.sayHello("Cok")
}
