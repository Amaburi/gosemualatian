package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk memeriksa apakah karakter yang diinputkan termasuk huruf hidup (vokal) atau bukan.

// aiueo

// Masukan berupa sebuah huruf (karakter).

// Keluaran berupa boolean true jika karakter termasuk huruf vokal atau false jika bukan.

// Catatan: Masukan selalu dalam bentuk huruf kecil.
func main() {
	var inp byte
	var hasil bool
	fmt.Scanf("%c", &inp)
	hasil = inp == 'a' || inp == 'i' || inp == 'u' || inp == 'e' || inp == 'o'

	fmt.Print(hasil)
}
