package main

import (
	"fmt"
	"math"
)

// Masukan berupa bilangan real yang menyatakan luas lingkaran.
// Keluaran berupa diamater lingkaran dalam bilangan real.
// Petunjuk: Diameter lingkaran adalah 2 x panjang jari-jari. Gunakan pi = 3.14

func main() {
	var luas float64
	var diameter float64
	fmt.Scan(&luas)
	diameter = 2 * math.Sqrt(luas/3.14)
	fmt.Println(diameter)
}
