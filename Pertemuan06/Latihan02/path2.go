package main

import (
	"fmt"
	"os"
)

func main() {
	var err error
	err = os.Chmod("RAMADHANI", 1019)
	if err != nil{
		fmt.Println("Error:", err )
		return
	}
	fmt.Println("Izin 'RAMADHANI'telah diubah menjadi 1019")
}