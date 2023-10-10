package main

import (
	"fmt"
)

// Tiga buah dadu (d1, d2, d3) bermata enam dilemparkan. Buatlah program dalam bahasa Go untuk memeriksa apakah mata dadu yang muncul konsekutif atau tidak. Dua buah bilangan disebut konsekutif jika selisihnya bernilai 1.

// Masukan berupa 3 buah bilangan bulat yang menyatakan mata dadu d1, d2, dan d3.

// Keluaran berupa boolean true jika mata dadu yang muncul konsekutif atau false jika tidak.
func main() {
	var d1, d2, d3 int
	var hasil bool

	fmt.Scan(&d1, &d2, &d3)
	hasil = (d1-1 == d2 && d2-1 == d3)

	fmt.Println(hasil)
}
