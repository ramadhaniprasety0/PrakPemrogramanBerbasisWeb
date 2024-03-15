package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Membuat scanner untuk input dari konsol
	scanner := bufio.NewScanner(os.Stdin)

	// Input nama
	fmt.Print("Masukkan nama Anda: ")
	scanner.Scan()
	nama := scanner.Text()

	// Input usia
	fmt.Print("Masukkan usia Anda: ")
	scanner.Scan()
	var usia int
	fmt.Sscanf(scanner.Text(), "%d", &usia)

	// Menentukan kategori usia
	var kategori string
	if usia < 18 {
		kategori = "anak-anak"
	} else {
		kategori = "dewasa"
	}

	// Menampilkan pesan
	fmt.Printf("Selamat datang, %s! Anda termasuk kategori %s.\n", nama, kategori)
}
