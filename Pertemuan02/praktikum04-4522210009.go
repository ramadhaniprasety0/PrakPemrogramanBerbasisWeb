package main

import "fmt"

func main() {
	panjang := 5
	lebar := 3

	luas, keliling := hitungLuasKelilingPersegiPanjang(panjang, lebar)

	fmt.Println("Luas persegi panjang:", luas)
	fmt.Println("Keliling persegi panjang:", keliling)
}

func hitungLuasKelilingPersegiPanjang(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)
	return
}
