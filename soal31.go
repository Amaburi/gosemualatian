package main

import (
	"fmt"
)

func main() {
	var inp, inp2 byte
	var hasil bool

	fmt.Scanf("%c", &inp, &inp2)
	hasil = inp == (inp2-32) || (inp-32) == inp2
	fmt.Printf("%c", hasil)

}
