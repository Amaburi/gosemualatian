package main

import (
	"fmt"
)

// Tiga buah dadu (d1, d2, d3) bermata enam dilemparkan. Buatlah program dalam bahasa Go untuk memeriksa apakah ada mata dadu ganjil yang muncul atau tidak.

// Masukan berupa 3 buah bilangan bulat yang menyatakan mata dadu d1, d2, dan d3.

// Keluaran berupa boolean true jika terdapat mata dadu ganjil muncul atau false jika tidak.
func main() {
	var d1, d2, d3 int
	var hasil bool

	fmt.Scan(&d1, &d2, &d3)
	hasil = d1%2 == 1 || d2%2 == 1 || d3%2 == 1
	fmt.Println(hasil)
}
