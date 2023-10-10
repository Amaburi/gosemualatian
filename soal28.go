package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk mengkonversi huruf besar (uppercase) ke huruf kecil (lowercase).

// Masukan berupa sebuah karakter yang merupakan huruf besar dalam alfabet (uppercase).

// Keluaran berupa sebuah karakter yang menyatakan huruf kecil dalam alfabet (lowercase).
func main() {
	var inp byte
	fmt.Scanf("%c", &inp)
	fmt.Printf("%c", inp+32)
}
