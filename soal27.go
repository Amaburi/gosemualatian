package main

import (
	"fmt"
)

// Setiap pelanggan yang berbelanja di sebuah swalayan mendapatkan kode berupa bilangan bulat dengan banyak digit sebanyak 4 buah. Cashback akan diberikan jika digit pertama bilangan itu sama dengan digit terakhir. Buatlah program dalam bahasa Go untuk memeriksa apakah seorang pelanggan akan mendapatkan cashback atau tidak.

// Masukan berupa bilangan bulat berjumlah digit 4.

// Keluaran berupa boolean true jika mendapat cashback, atu false jika tidak.
func main() {
	var inp int
	var cashback bool
	fmt.Scanln(&inp)
	stringinp := fmt.Sprintf("%04d", inp)
	cashback = stringinp[0] == stringinp[3]
	fmt.Println(cashback)
}
func caralain() {
	var inp int
	var cashback bool
	fmt.Scanln(&inp)
	digitsatu := inp / 1000
	digitakhir := inp % 10
	cashback = digitsatu == digitakhir
	fmt.Println(cashback)
}
