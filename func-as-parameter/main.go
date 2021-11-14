package main

import (
	"fmt"
)

func sayHello(name string, filter func(string) string) {
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
