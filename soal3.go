package main

import "fmt"

// Buatlah program dalam bahasa Go untuk menghitung nilai fungsi f(x)=x3+3xx4−3x2+4

// Masukan berupa bilangan real yang menyatakan x.

// Keluaran berupa nilai fungsi f(x).

func main() {
	var x, fungsi float64
	fmt.Scan(&x)
	fungsi = (x*x*x + 3*x) / (x*x*x*x - 3*x*x + 4)
	fmt.Println(fungsi)
}
