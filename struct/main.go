package main

import "fmt"

type Costumer struct {
	name, adress string
	age, number  int
}

func main() {
	// saepudin := Costumer{
	// 	name:   "Uhuyyy",
	// 	adress: "sdaadad",
	// 	age:    21,
	// 	number: 2312312321,
	// }

	saepudin := Costumer{"Uhuuuyy", "sdasdasd", 21, 231231}
	fmt.Println(saepudin)
}
