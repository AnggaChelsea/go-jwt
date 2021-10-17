package main

import "fmt"

type Customer struct {
	City, Address string
	Zip           int
}

func main() {
	address := Customer{"Sukabumi", "Kp,Gobang", 4322}
	address2 := &address
	address2.City = "Jakarta"
	*address2 = Customer{"Sedney", "Kp,Gobang", 4322}

	fmt.Println(address)
	fmt.Println(address2)

	
}
