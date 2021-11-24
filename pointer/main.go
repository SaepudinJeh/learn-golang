package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Cirebon", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Majalengka"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
}
