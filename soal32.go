package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go suatu aplikasi kalkulator yang digunakan untuk menghitung operasi aritmatika pada bilangan pecahan a/b dan c/d.

// Masukan terdiri dari 4 buah bilangan bulat secara berurutan: a, b, c, dan d; yang digunakan untuk membentuk pecahan a/b dan c/d.

// Keluaran terdiri 4 baris, di mana masing-masing barisnya terdapat 2 buah bilangan bulat x dan y, yang membentuk pecahan x/y hasil operasi aritmatika. Masing-masing baris adalah hasil operasi penjumlahan, pengurangan, perkalian dan pembagian.

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	// Penjumlahan
	numerator1 := a*d + b*c
	denominator1 := b * d

	// Pengurangan
	numerator2 := a*d - b*c
	denominator2 := b * d

	// Perkalian
	numerator3 := a * c
	denominator3 := b * d

	// Pembagian
	numerator4 := a * d
	denominator4 := b * c
	fmt.Printf("penjumlahan: %d / %d\n", numerator1, denominator1)
	fmt.Printf("pengurangan: %d / %d\n", numerator2, denominator2)
	fmt.Printf("perkalian: %d / %d\n", numerator3, denominator3)
	fmt.Printf("pembagian: %d / %d\n", numerator4, denominator4)
}
