package main

import (
	"fmt"
)

// Adik dan kakak sedang bermain tebak-tebakan bilangan. Keduanya menuliskan sebuah bilangan dari 0 hingga 9 di atas secarik kertas secara rahasia dan kemudian membandingkannya. Adik dinyatakan pemenang jika bilangan adik 2x dari bilangan kakak atau bilangan adik 3x dari bilangan kakak. Buatlah program dalam bahasa Go untuk menentukan apakah adik pemenang permainan ini atau bukan.

// Masukan berupa dua bilangan bulat. Bilangan pertama adalah bilangan tebakan adik, sedangkan bilangan kedua adalah bilangan tebakan kakak.

// Keluaran berupa boolean true atau false. True berarti adik pemenang permainan ini.
func main() {
	var x, y int
	var hasil bool
	fmt.Scan(&x, &y)
	hasil = (x == 2*y) || (x == 3*y)
	fmt.Print(hasil)
}
