package main

import (
	"fmt"
)

// Adik dan kakak sedang bermain tebak-tebakan bilangan. Keduanya menuliskan sebuah bilangan dari 0 hingga 9 di atas secarik kertas secara rahasia dan kemudian membandingkannya. Adik dinyatakan pemenang jika bilangannya sama dengan bilangan kakak atau selisih bilangannya 3 atau 6. Buatlah program dalam bahasa Go untuk menentukan apakah adik pemenang permainan ini.

// Masukan berupa dua bilangan bulat. Bilangan pertama adalah bilangan tebakan adik, sedangkan bilangan kedua adalah bilangan tebakan kakak.

// Keluaran berupa boolean true atau false. True berarti adik pemenang permainan ini.

func main() {
	var x, y int
	var hasil bool
	fmt.Scan(&x, &y)
	hasil = (x == y) || (x-y)%3 == 0 || (x-y)%6 == 0
	fmt.Print(hasil)
}
