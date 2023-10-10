package main

import (
	"fmt"
)

func main() {
	var inp, jam, menit, detik, sisa int
	fmt.Scanln(&inp)
	jam = inp / 3600
	sisa = inp % 3600
	menit = sisa / 60
	detik = inp % 60
	fmt.Println(jam, menit, detik)

}
