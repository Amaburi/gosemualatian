package main

import "fmt"

// Adik dan kakak sedang bermain tebak-tebakan bilangan. Keduanya menuliskan sebuah bilangan dari 0 hingga 9 di atas secarik kertas secara rahasia dan kemudian membandingkannya. Adik dinyatakan pemenang jika bilangannya sama dengan bilangan kakak atau selisih bilangannya 1 atau 5. Buatlah algoritma untuk menentukan apakah adik pemenang permainan ini.

// Masukan berupa dua bilangan bulat. Bilangan pertama adalah bilangan tebakan adik, sedangkan bilangan kedua adalah bilangan tebakan kakak.

// Keluaran berupa boolean true atau false. True berarti adik pemenang permainan ini.
func main() {
	var x1, x2 int
	var hasil bool

	fmt.Scan(&x1, &x2)
	hasil = (x1 == x2) || (x1-x2 == 1) || (x1-x2 == -1) || (x1-x2 == 5) || (x1-x2 == -5)

	fmt.Print(hasil)
}
