package main

import (
	"fmt"
)

// Setiap pelanggan yang berbelanja di sebuah swalayan mendapatkan cashback jika mempunyai kartu member dan total belanjanya minimal Rp 1 juta. Buatlah program dalam bahasa Go untuk memeriksa apakah seorang pelanggan akan mendapatkan cashback atau tidak.

// Masukan berupa boolean dan bilangan bulat yang menyatakan mempunyai/tidak mempunyai kartu (true/false) dan total belanja.

// Keluaran berupa boolean true jika mendapat cashback, atu false jika tidak.
func main() {
	var uang int
	var kartu, cashback bool
	fmt.Scan(&kartu, &uang)
	cashback = kartu == true && uang >= 1000000
	fmt.Println(cashback)
}
