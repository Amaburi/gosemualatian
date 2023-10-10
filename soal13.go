package main

import (
	"fmt"
)

func main() {
	var x, y int
	var hasil bool
	fmt.Scan(&x, &y)
	hasil = (x-y)%2 == 0
	fmt.Print(hasil)
}
