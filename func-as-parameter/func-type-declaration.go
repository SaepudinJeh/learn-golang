package main

import (
	"fmt"
)

type Filter func(string) string

func sayHello(name string, filter Filter) {
	nameFilter := filter(name)
	fmt.Println("hello", nameFilter)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHello("Saepudin", spamFilter)
	sayHello("Anjing", spamFilter)
}
