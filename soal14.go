package main

import (
	"fmt"
)

// Adik dinyatakan pemenang jika selisih bilangan adik dengan bilangan kakak adalah bilangan prima

// Masukan berupa dua bilangan bulat. Bilangan pertama adalah bilangan tebakan adik, sedangkan bilangan kedua adalah bilangan tebakan kakak.

// Keluaran berupa boolean true atau false. True berarti adik pemenang permainan ini.
func main() {
	var x, y int
	var hasil bool
	fmt.Scan(&x, &y)
	hasil = (x-y)%2 == 0 || (x-y)%3 == 0 || (x-y)%5 == 0 || (x-y)%7 == 0
	fmt.Print(hasil)
}
