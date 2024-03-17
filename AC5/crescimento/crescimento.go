package main

import (
	"fmt"
)

func main() {
	var T, PA, PB int
	var G1, G2 float32
	fmt.Scanln(&T)

	for T != 0 {
		fmt.Scanln(&PA, &PB, &G1, &G2)
		i := 0
		for PA <= PB {
			PA = int(float32(PA) * (1 + G1/100))
			PB = int(float32(PB) * (1 + G2/100))
			i++
			if i > 100 {
				fmt.Println("Mais de 1 seculo.")
				break
			}
		}

		if i <= 100 {
			fmt.Println(i, "anos.")
		}
		T--
	}
}
