package main

import "fmt"

// Buatlah program dalam bahasa Go untuk memeriksa apakah huruf yang diinputkan termasuk huruf mati (konsonan) atau bukan.

// all konsonan  words

// Masukan berupa sebuah huruf (karakter).

// Keluaran berupa boolean true jika huruf termasuk konsonan atau false jika bukan.

// Catatan: Masukan selalu dalam huruf kecil.
func main() {
	var inp byte
	var hasil bool
	fmt.Scanf("%c", &inp)

	hasil = inp != 'a' && inp != 'i' && inp != 'u' && inp != 'e' && inp != 'o'

	fmt.Print(hasil)
}
