package main

import (
	"fmt"
	"math"
)

func main() {
	var N, H, C, L float64
	for {
		_, err1 := fmt.Scanln(&N)
		if err1 != nil {
			return
		}
		fmt.Scan(&H, &C, &L)
		resultado := math.Sqrt(math.Pow(N*C, 2)+math.Pow(N*H, 2)) * L / 10000.0
		fmt.Printf("%.4f\n", resultado)
	}
}
