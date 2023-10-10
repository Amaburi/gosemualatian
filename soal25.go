package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk mengkonversi waktu dalam bilangan bulat dengan satuan jam, menit, dan detik menjadi satuan detik.
// Masukan berupa 3 bilangan bulat yang menyatakan waktu dalam jam, menit, dan detik.

// Keluaran berupa bilangan bulat yang menyatakan hasil konversi ke detik.
func main() {
	var inp, inp2, inp3, jam, menit, detik, hasil int
	fmt.Scan(&inp, &inp2, &inp3)
	jam = inp * 3600
	menit = inp2 * 60
	detik = inp3
	hasil = jam + menit + detik
	fmt.Println(hasil)

}
