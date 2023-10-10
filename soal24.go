package main

import (
	"fmt"
)

// Buatlah program dalam bahasa Go untuk menghitung persentase kelulusan dan ketidaklulusan mahasiswa. Misalkan masukan 12 8, maka keluarannya 0.60 0.40.

// Masukan berupa dua bilangan bulat yang menyatakan jumlah mahasiswa yang lulus dan tidak lulus.

// Keluaran berupa bilangan yang menyatakan persentase kelulusan dan ketidaklulusan.

// Catatan: Bulatkan 2 digit di belakang koma.
func main() {
	var inp1, inp2, totalmahasiswa int
	var lulus, tidaklulus float64
	fmt.Scan(&inp1, &inp2)
	totalmahasiswa = inp1 + inp2
	lulus = float64(inp1) / float64(totalmahasiswa)
	tidaklulus = float64(inp2) / float64(totalmahasiswa)
	fmt.Printf("%.2f %.2f\n", lulus, tidaklulus)
}
