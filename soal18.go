package main

import (
	"fmt"
)

// Tiga buah dadu (d1, d2, d3) bermata enam dilemparkan. Buatlah program dalam bahasa Go untuk memeriksa apakah mata dadu yang muncul semuanya ganjil atau tidak.

// Masukan berupa 3 buah bilangan bulat yang menyatakan mata dadu d1, d2, dan d3.

// Keluaran berupa boolean true jika ketiga mata dadu yang muncul semuanya ganjil atau false jika tidak.
func main() {
	var d1, d2, d3 int
	var hasil bool

	fmt.Scan(&d1, &d2, &d3)
	hasil = (d1+d2+d3) == 3 || (d1+d2+d3) == 5 || (d1+d2+d3) == 7 || (d1+d2+d3) == 9
	fmt.Println(hasil)

}
func anothertrick() {
	var d1, d2, d3 int
	var hasil bool

	fmt.Scan(&d1, &d2, &d3)
	hasil = d1%2 == 1 && d2%2 == 1 && d3%2 == 1
	fmt.Println(hasil)
}
