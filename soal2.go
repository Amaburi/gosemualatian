package main

import (
	"fmt"
	"math"
)

// Buatlah program dalam bahasa Go untuk menghitung  y = (x**2+1/x)2
// .

// Masukan berupa sebuah bilangan real yang menyatakan x.

// Keluaran berupa nilai y.
func main() {
	var x, y float64
	fmt.Scan(&x)
	y = math.Pow(x*x+1/x, 2)
	fmt.Println(y)
}
