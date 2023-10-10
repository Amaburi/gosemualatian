package main

import (
	"fmt"
)

func main() {
	var inp1 int
	fmt.Scan(&inp1)
	d1 := inp1 / 1000
	d2 := (inp1 / 100) % 10
	d3 := (inp1 / 10) % 10
	d4 := inp1 % 10
	hasil := (d1 <= d2) && (d2 <= d3) && (d3 <= d4)
	fmt.Println(hasil)
}
