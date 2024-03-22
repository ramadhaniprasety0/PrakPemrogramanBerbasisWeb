package main

import "fmt"

func main() {
	batasNilai := map[string]int{
		"A": 90,
		"B": 80,
		"C": 70,
		"D": 60,
	}

	getNilaiHuruf := func(nilaiUjian int) string {
		for hufuf, batas := range batasNilai {
			if nilaiUjian >= batas {
				return hufuf
			}
		}
		return "E"
	}

	nilaiUjian := []int{85, 75, 95, 55, 65}
	for _, nilai := range nilaiUjian {
		fmt.Println("Nilai Ujian ", nilai, ":", getNilaiHuruf(nilai))
	}

}