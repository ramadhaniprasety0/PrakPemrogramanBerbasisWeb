package main

import (
	"Praktikum_03/ratarata"
	"fmt"
)

func main()  {
	nilaiSiswa := []int{80, 75, 90, 85, 60}

	rataRata := ratarata.HitungRataRata(nilaiSiswa)

	fmt.Println("Rata-Rata nilai siswa : ", rataRata)
}