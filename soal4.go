package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk mencetak karakter yang dimasukkan.

// Masukan berupa sebuah karakter.

// Keluaran berupa sebuah karakter sebagaimana contoh.

// Petunjuk: Gunakan fmt.Scanf untuk pembacaan masukan dan fmt.Printf untuk pencetakan.
func main() {
	var inp byte
	fmt.Scanf("%c", &inp)
	fmt.Printf("%c", inp)
}
