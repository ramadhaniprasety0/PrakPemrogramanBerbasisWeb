package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Mahasiswa struct {
	Nama    string
	NPM     string
	Jurusan string
}

func main() {

	dataMahasiswa := make(map[string]Mahasiswa)

	for {
		mahasiswaBaru := inputMahasiswa()
		dataMahasiswa[mahasiswaBaru.NPM] = mahasiswaBaru

		fmt.Print("Tambah data mahasiswa lagi? (y/n): ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		jawaban := scanner.Text()
		if strings.ToLower(jawaban) != "y" {
			break
		}
	}

	fmt.Println("\nDaftar Mahasiswa:")
	for _, mhs := range dataMahasiswa {
		fmt.Println(mhs.Nama)
	}

	fmt.Print("\nMasukkan NPM mahasiswa yang ingin dicari: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	npmCari := scanner.Text()
	mahasiswaCari, found := dataMahasiswa[npmCari]
	if found {
		fmt.Println("\nData Mahasiswa:")
		fmt.Println("Nama:", mahasiswaCari.Nama)
		fmt.Println("NPM:", mahasiswaCari.NPM)
		fmt.Println("Jurusan:", mahasiswaCari.Jurusan)
	} else {
		fmt.Println("Mahasiswa dengan NPM tersebut tidak ditemukan.")
	}
}

func inputMahasiswa() Mahasiswa {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan Nama Mahasiswa: ")
	scanner.Scan()
	nama := scanner.Text()

	fmt.Print("Masukkan NPM Mahasiswa: ")
	scanner.Scan()
	npm := scanner.Text()

	fmt.Print("Masukkan Jurusan Mahasiswa: ")
	scanner.Scan()
	jurusan := scanner.Text()

	return Mahasiswa{Nama: nama, NPM: npm, Jurusan: jurusan}
}
