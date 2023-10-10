package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk menentukan apakah tahun yang dimasukkan termasuk kebisat atau bukan.

// Masukan berupa bilangan bulat positif yang menyatakan tahun.

// Keluaran berupa boolean true jika kabisat, atau false jika bukan.
func main() {
	var inp int
	var hasil bool
	fmt.Scan(&inp)
	hasil = (inp%4 == 0 && inp%100 != 0) || (inp%400 == 0)

	fmt.Print(hasil)
}
