package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk mempertukarkan nilai bilangan bulat a, b, c, d dengan ketentuan sebagai berikut:

// nilai a berisi nilai d
// nilai d berisi nilai c
// nilai c berisi nilai b
// nilai b berisi nilai a
// Masukan berupa 4 buah bilangan bulat a, b, c, d.

// Keluaran berupa nilai a, b, c, d.
func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	a, b, c, d = b, c, d, a
	fmt.Println(a, b, c, d)
}
