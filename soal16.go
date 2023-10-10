package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk mempertukarkan nilai bilangan bulat x, y, z dengan ketentuan sebagai berikut:

// nilai y berisi nilai x
// nilai x berisi nilai z
// nilai z berisi nilai y
// Masukan berupa 3 buah bilangan bulat x, y, z.

// Keluaran berupa nilai x, y, z setelah pertukaran.
func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	x, y, z = z, x, y

	fmt.Println(x, y, z)
}
