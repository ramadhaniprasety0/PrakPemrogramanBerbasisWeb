package main

import "fmt"

func main() {
	panjang := 5
	lebar := 3

	luas := hitungLuasPersegiPanjang(panjang, lebar)
	keliling := hitungKelilingPersegiPanjang(panjang, lebar)

	fmt.Println("Luas persegi panjang:", luas)
	fmt.Println("Keliling persegi panjang:", keliling)
}

func hitungLuasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func hitungKelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}
