package main

import (
	"fmt"
	"math"
)

func main() {
	jariJari := 5.0

	luas := func(r float64) float64 {
		return math.Pi * r * r
	}(jariJari)

	keliling := func (r float64) float64  {
		return 2 * math.Pi * r
	}(jariJari)

	fmt.Println("Luas Lingkaran : ", luas)
	fmt.Println("Keliling Linkaran : ", keliling)
}