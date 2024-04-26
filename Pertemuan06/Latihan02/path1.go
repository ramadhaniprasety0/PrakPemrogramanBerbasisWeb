package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("RAMADHANI", 1009)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Direktori 'RAMADHANI'berhasil dibuat. ")
}